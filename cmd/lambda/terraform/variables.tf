variable "aws_region" {
  description = "AWS region for Lambda deployment"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment name (local/prod)"
  type        = string
  default     = "local"

  validation {
    condition     = contains(["local", "prod"], var.environment)
    error_message = "Environment must be local or prod."
  }
}

variable "lambda_timeout" {
  description = "Lambda function timeout in seconds"
  type        = number
  default     = 30

  validation {
    condition     = var.lambda_timeout >= 1 && var.lambda_timeout <= 900
    error_message = "Timeout must be between 1 and 900 seconds."
  }
}

variable "lambda_memory_size" {
  description = "Lambda memory in MB (affects CPU allocation)"
  type        = number
  default     = 512

  validation {
    condition     = var.lambda_memory_size >= 128 && var.lambda_memory_size <= 10240
    error_message = "Memory must be between 128 and 10240 MB."
  }
}

variable "api_gateway_cors_allowed_origins" {
  description = "Allowed origins for CORS (use ['*'] for public access)"
  type        = list(string)
  default     = ["*"]
}

variable "api_gateway_cors_allowed_methods" {
  description = "Allowed HTTP methods for CORS"
  type        = list(string)
  default     = ["POST", "OPTIONS"]
}

variable "api_gateway_cors_max_age" {
  description = "Max age for CORS preflight cache (seconds)"
  type        = number
  default     = 86400 # 24 hours
}
