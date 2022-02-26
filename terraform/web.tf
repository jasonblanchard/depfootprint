resource "aws_s3_bucket" "web" {
  bucket = "depfootprint-static-assets"
}

resource "aws_s3_bucket_website_configuration" "web" {
  bucket = aws_s3_bucket.web.bucket

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }
}

resource "aws_apigatewayv2_integration" "web" {
  api_id             = aws_apigatewayv2_api.gw.id
  integration_type   = "HTTP_PROXY"
  integration_uri    = "http://${aws_s3_bucket.web.website_endpoint}"
  integration_method = "GET"
}

resource "aws_apigatewayv2_route" "web" {
  api_id    = aws_apigatewayv2_api.gw.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.web.id}"
}

output "web_host" {
  value = aws_s3_bucket.web.website_endpoint
}