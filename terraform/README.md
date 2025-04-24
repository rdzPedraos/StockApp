# Infraestructura de Stock App con Terraform

Este directorio contiene la configuración de Terraform para desplegar la aplicación Stock App en AWS.

## Componentes desplegados

- **VPC** con 2 subredes públicas
- **ECS Cluster** con 2 servicios:
  - API (Go)
  - Cliente (Nuxt.js)
- **ElastiCache** para Redis
- **Application Load Balancers** para la API y el Cliente
- **Grupos de seguridad** para controlar el tráfico entre componentes

## Requisitos previos

1. Terraform instalado (versión >= 1.0.0)
2. AWS CLI configurado con credenciales adecuadas
3. Imágenes de Docker para la API y el Cliente subidas a ECR

## Configuración

Antes de aplicar la configuración, necesitas crear un bucket S3 para el estado de Terraform:

```bash
aws s3 mb s3://stock-app-terraform-state --region us-east-1
```

Luego, crea un archivo `terraform.tfvars` con tus valores específicos:

```hcl
aws_region         = "us-east-1"  # O la región que prefieras
ecr_repository_url = "012345678901.dkr.ecr.us-east-1.amazonaws.com"  # Tu cuenta de AWS + región
app_name           = "stock-app"
```

## Despliegue

1. Inicializa Terraform:

```bash
terraform init
```

2. Verifica los cambios que se aplicarán:

```bash
terraform plan
```

3. Aplica la configuración:

```bash
terraform apply
```

4. Cuando todo esté listo, recibirás las URLs para acceder a la aplicación.

## Preparación de imágenes para ECR

Para subir las imágenes Docker a ECR, sigue estos pasos:

1. Construye las imágenes de Docker:

```bash
docker-compose build
```

2. Crea repositorios ECR para la API y el Cliente:

```bash
aws ecr create-repository --repository-name stock-api
aws ecr create-repository --repository-name stock-client
```

3. Etiqueta y sube las imágenes:

```bash
# Obtén el login para ECR
aws ecr get-login-password | docker login --username AWS --password-stdin CUENTA.dkr.ecr.REGION.amazonaws.com

# Etiqueta las imágenes
docker tag stock-api:latest CUENTA.dkr.ecr.REGION.amazonaws.com/stock-api:latest
docker tag stock-client:latest CUENTA.dkr.ecr.REGION.amazonaws.com/stock-client:latest

# Sube las imágenes
docker push CUENTA.dkr.ecr.REGION.amazonaws.com/stock-api:latest
docker push CUENTA.dkr.ecr.REGION.amazonaws.com/stock-client:latest
```

## Limpieza

Para eliminar todos los recursos creados:

```bash
terraform destroy
```
