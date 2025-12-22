locals {
  project_name = "typedef"

  lambda_function_name  = "${local.project_name}-codegen-${var.environment}"
  lambda_role_name      = "${local.project_name}-lambda-role-${var.environment}"
  lambda_log_group_name = "/aws/lambda/${local.lambda_function_name}"
  lambda_timeout        = var.lambda_timeout
  lambda_memory_size    = var.lambda_memory_size
  lambda_concurrent_executions = var.lambda_concurrent_executions

  api_gateway_name                 = "${local.project_name}-api-${var.environment}"
  api_gateway_stage_name           = var.environment
  api_gateway_rate_limit           = var.api_gateway_rate_limit
  api_gateway_burst_limit          = var.api_gateway_burst_limit
  api_gateway_cors_allowed_origins = var.api_gateway_cors_allowed_origins
  api_gateway_cors_allowed_methods = var.api_gateway_cors_allowed_methods
  api_gateway_cors_max_age         = var.api_gateway_cors_max_age

  common_tags = {
    Project     = local.project_name
    Environment = var.environment
    Repository  = "github.com/exanubes/typedef"
  }

  lambda_runtime      = "provided.al2023"
  lambda_handler      = "bootstrap"
  lambda_architecture = ["arm64"]
}
