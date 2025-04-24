# StockApp - Aplicación de Gestión de Inventario

## Descripción

StockApp es una aplicación completa de gestión de inventario desarrollada como prueba técnica. Está diseñada con una arquitectura de microservicios que permite una gestión eficiente y escalable de productos, pedidos y gestión de stock.

## Stack Tecnológico

### Backend (API)

- **Lenguaje**: Go 1.24
- **Framework**: Gin para API REST
- **ORM**: GORM con PostgreSQL
- **Caché**: Redis para almacenamiento en caché y reducción de peticiones repetitivas
- **Comandos**: Cobra para CLI
- **Concurrencia**: Uso de goroutines para operaciones asíncronas y procesamiento paralelo

### Frontend (Cliente)

- **Framework**: Nuxt.js 3 (Vue.js)
- **UI**: Nuxt UI con Tailwind CSS
- **Gestión de estado**: Pinia
- **Peticiones HTTP**: Axios

### Infraestructura

- **Containerización**: Docker y Docker Compose
- **IaC**: Terraform para desplegar en AWS
- **Servicios AWS**: ECS (Fargate), ElastiCache (Redis), Application Load Balancer

## Características Principales

- **Diseño RESTful**: API con endpoints estandarizados siguiendo convenciones REST
- **Recursos estandarizados**: Implementación de un patrón de recursos uniforme para todas las entidades
- **Caché inteligente**: Sistema de caché para reducir peticiones repetitivas y mejorar el rendimiento
- **Concurrencia**: Uso de goroutines para operaciones asíncronas
- **UI Moderna**: Interfaz de usuario intuitiva y responsiva
- **Despliegue automatizado**: Infraestructura como código con Terraform
- **Containers**: Dockerfiles optimizados para desarrollo y producción

## Estructura del Proyecto

```
/
├── api/               # Aplicación backend en Go
├── client/            # Aplicación frontend en Nuxt.js
├── terraform/         # Configuración de infraestructura con Terraform
├── docker-compose.yml # Configuración para desarrollo local
└── README.md          # Este archivo
```

## Documentación

Para más detalles sobre cada componente del proyecto, consulta la documentación específica:

- [Documentación de la API](./api/readme.md)
- [Documentación del Cliente](./client/README.md)
- [Documentación de Terraform](./terraform/README.md)

## Desarrollo Local

Para ejecutar el proyecto en tu entorno local:

```bash
# Construir y levantar todos los servicios
docker-compose up -d

# La API estará disponible en: http://localhost:8080
# El cliente estará disponible en: http://localhost:3000
```

## Despliegue en AWS

El proyecto incluye configuración completa para desplegar en AWS usando Terraform:

```bash
cd terraform
terraform init
terraform apply
```

Para más detalles sobre el despliegue, consulta la [documentación de Terraform](./terraform/README.md).

## Optimizaciones

- **Caché Redis**: Implementación de caché para reducir peticiones repetitivas a la base de datos
- **Multi-stage Docker builds**: Reducción del tamaño de las imágenes para producción
- **Escalabilidad**: Configuración de infraestructura para escalar horizontalmente
- **Concurrencia**: Uso de goroutines para operaciones paralelas y mejor rendimiento

## Estándar de Recursos

La API implementa un patrón uniforme para todos los recursos:

- **GET /resource**: Listar recursos
- **GET /resource/:id**: Obtener un recurso específico
- **POST /resource**: Crear un nuevo recurso
- **PUT /resource/:id**: Actualizar un recurso existente
- **DELETE /resource/:id**: Eliminar un recurso

Este estándar facilita el consumo y mantenimiento de la API.
