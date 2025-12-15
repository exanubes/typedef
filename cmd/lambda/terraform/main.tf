resource "aws_iam_role" "lambda_role" {
  name = local.lambda_role_name

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

  tags = local.common_tags
}

resource "aws_iam_role_policy_attachment" "lambda_basic_policy" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}


resource "aws_cloudwatch_log_group" "lambda_logs" {
  name              = local.lambda_log_group_name
  retention_in_days = 1 

  tags = local.common_tags
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_dir  = "${path.module}/../dist/lambda"
  output_path = "${path.module}/../dist/terraform/typedef-lambda.zip"
}

resource "aws_lambda_function" "typedef" {
  filename         = data.archive_file.lambda_zip.output_path
  function_name    = local.lambda_function_name
  role            = aws_iam_role.lambda_role.arn
  handler         = local.lambda_handler
  runtime         = local.lambda_runtime
  architectures   = local.lambda_architecture

  source_code_hash = data.archive_file.lambda_zip.output_base64sha256

  timeout     = var.lambda_timeout
  memory_size = var.lambda_memory_size

  depends_on = [
    aws_cloudwatch_log_group.lambda_logs,
    aws_iam_role_policy_attachment.lambda_basic_policy
  ]

  tags = local.common_tags
}

resource "aws_apigatewayv2_api" "typedef" {
  name          = local.api_gateway_name
  protocol_type = "HTTP"

  cors_configuration {
    allow_origins     = var.api_gateway_cors_allowed_origins
    allow_methods     = var.api_gateway_cors_allowed_methods
    allow_headers     = ["content-type"]
    expose_headers    = ["content-type", "x-amz-request-id"]
    max_age          = var.api_gateway_cors_max_age
    allow_credentials = false
  }

  tags = local.common_tags
}

resource "aws_apigatewayv2_stage" "typedef" {
  api_id      = aws_apigatewayv2_api.typedef.id
  name        = local.api_gateway_stage_name
  auto_deploy = true

  tags = local.common_tags
}

resource "aws_apigatewayv2_integration" "typedef" {
  api_id             = aws_apigatewayv2_api.typedef.id
  integration_type   = "AWS_PROXY"
  integration_uri    = aws_lambda_function.typedef.invoke_arn
  integration_method = "POST"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "typedef" {
  api_id    = aws_apigatewayv2_api.typedef.id
  route_key = "POST /codegen"
  target    = "integrations/${aws_apigatewayv2_integration.typedef.id}"
}

resource "aws_lambda_permission" "api_gateway" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.typedef.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.typedef.execution_arn}/*/*"
}
