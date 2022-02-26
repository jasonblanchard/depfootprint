resource "aws_s3_bucket" "api" {
  bucket = "depfootprint-api-source"
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "depfootprint-lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "api" {
  function_name    = "depfootprint-api"
  s3_bucket        = aws_s3_bucket.api.bucket
  s3_key           = "api.zip"
  handler          = "api"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "go1.x"
  publish          = true
  source_code_hash = filebase64sha256("../bin/lambda/api.zip")
  timeout          = 120
}

resource "aws_lambda_permission" "lambda_permission" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.api.function_name
  principal     = "apigateway.amazonaws.com"

  # The /*/*/* part allows invocation from any stage, method and resource path
  # within API Gateway REST API.
  source_arn = "${aws_apigatewayv2_api.gw.execution_arn}/*/*/*"
}

resource "aws_apigatewayv2_integration" "api" {
  api_id                 = aws_apigatewayv2_api.gw.id
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.api.invoke_arn
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "api" {
  api_id    = aws_apigatewayv2_api.gw.id
  route_key = "ANY /api/{proxy+}"

  target = "integrations/${aws_apigatewayv2_integration.api.id}"
}