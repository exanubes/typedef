locals {
  project_name = "typedef"

  lambda_function_name  = "${local.project_name}-codegen-${var.environment}"
  lambda_role_name      = "${local.project_name}-lambda-role-${var.environment}"
  lambda_log_group_name = "/aws/lambda/${local.lambda_function_name}"
  api_gateway_name      = "${local.project_name}-api-${var.environment}"
  api_gateway_stage_name = var.environment

  common_tags = {
    Project     = local.project_name
    Environment = var.environment
    Repository  = "github.com/exanubes/typedef"
  }

  lambda_runtime      = "provided.al2023"
  lambda_handler      = "bootstrap"
  lambda_architecture = ["arm64"]
}
