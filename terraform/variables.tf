variable "aws_region" {
  description = "Región de AWS donde se desplegarán los recursos"
  type        = string
  default     = "us-east-1"
}

variable "ecr_repository_url" {
  description = "URL base del repositorio ECR para las imágenes de contenedores"
  type        = string
}

variable "app_name" {
  description = "Nombre de la aplicación"
  type        = string
  default     = "stock-app"
} 