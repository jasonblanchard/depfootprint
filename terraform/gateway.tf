resource "aws_apigatewayv2_api" "gw" {
  name          = "depfootprint"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.gw.id
  name        = "$default"
  auto_deploy = true
}


output "gw_api_invoke_url" {
  value = "http://${aws_apigatewayv2_stage.default.invoke_url}"
}