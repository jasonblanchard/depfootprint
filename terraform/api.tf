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