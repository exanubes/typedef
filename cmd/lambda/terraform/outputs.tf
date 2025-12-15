output "lambda_function_arn" {
  description = "ARN of the Lambda function"
  value       = aws_lambda_function.typedef.arn
}

output "lambda_function_name" {
  description = "Name of the Lambda function"
  value       = aws_lambda_function.typedef.function_name
}

output "api_gateway_endpoint" {
  description = "HTTPS endpoint for the API Gateway (use this for API calls)"
  value       = aws_apigatewayv2_stage.typedef.invoke_url
}

output "api_gateway_id" {
  description = "API Gateway identifier"
  value       = aws_apigatewayv2_api.typedef.id
}

output "api_gateway_stage" {
  description = "API Gateway stage name"
  value       = aws_apigatewayv2_stage.typedef.name
}

output "lambda_role_arn" {
  description = "ARN of the Lambda execution role"
  value       = aws_iam_role.lambda_role.arn
}

output "cloudwatch_log_group" {
  description = "CloudWatch Logs group name for Lambda logs"
  value       = aws_cloudwatch_log_group.lambda_logs.name
}

output "curl_example" {
  description = "Example curl command to test the API"
  value       = <<-EOT
    curl -X POST ${aws_apigatewayv2_stage.typedef.invoke_url} \
      -H "Content-Type: application/json" \
      -d '{
        "input_type": "json",
        "data": "{\"name\":\"John\",\"age\":30}",
        "format": "go"
      }'
  EOT
}
