# Stock App API

## Descripción

Stock App API es un servicio backend desarrollado en Go que proporciona información financiera sobre acciones y tickers de bolsa. La aplicación permite consultar datos de empresas, obtener recomendaciones de inversión y acceder a información de mercado actualizada a través de integraciones con APIs externas.

## Tecnologías

- **Lenguaje**: Go
- **Framework Web**: Gin Gonic
- **Base de datos**: PostgreSQL/CockroachDB
- **Caché**: Redis
- **Integraciones**: APIs financieras, Gemini AI

## Estructura del Proyecto

```
├── cmd/                    # Comandos CLI para la aplicación
├── config/                 # Configuración de la aplicación
├── controllers/            # Controladores para las rutas HTTP
├── database/               # Configuración y conexión a la base de datos
├── enums/                  # Enumeraciones y constantes
├── integrations/           # Integraciones con servicios externos
│   ├── financial/          # Integración con APIs financieras
│   ├── gemini/             # Integración con Gemini AI
│   └── stockRating/        # Integración con servicios de calificación de acciones
├── lib/                    # Bibliotecas y utilidades
├── models/                 # Modelos de datos
├── resources/              # Patrón personalizado para transformar y formatear respuestas de API
├── routes/                 # Definición de rutas de la API
├── server/                 # Configuración del servidor
├── .env.example            # Ejemplo de variables de entorno
├── go.mod                  # Dependencias de Go
├── go.sum                  # Checksums de dependencias
└── main.go                 # Punto de entrada de la aplicación
```

## Modelos Principales

- **Ticker**: Representa una acción bursátil con su símbolo y nombre de empresa
- **Recommendation**: Representa recomendaciones de compra/venta para tickers
- **Broker**: Información sobre las casas de bolsa que emiten recomendaciones

## Rutas API

La API expone los siguientes endpoints:

- `GET /api/ping`: Verificación de estado de la API
- `GET /api/tickers`: Lista todos los tickers disponibles
- `GET /api/tickers/:id`: Obtiene información detallada sobre un ticker específico

## Integraciones

La aplicación se integra con varios servicios externos:

1. **APIs Financieras**: Para obtener datos de mercado, precios de acciones, capitalización bursátil y logotipos de empresas
2. **Gemini AI**: Para análisis y predicciones avanzadas
3. **Stock Rating API**: Para obtener calificaciones y recomendaciones de expertos

## Patrón Resources

El proyecto implementa un patrón personalizado de "Resources" que permite transformar y enriquecer las respuestas de los controladores antes de enviarlas al cliente. Este patrón:

- Encapsula la lógica de formateo y transformación de datos
- Permite enriquecer respuestas con datos de múltiples fuentes de forma concurrente
- Facilita la consistencia en el formato de las respuestas API
- Implementa una interfaz genérica `TypedResource` para diferentes tipos de datos

## Configuración

### Variables de Entorno

Copia el archivo `.env.example` a `.env` y configura las siguientes variables:

```
# Configuración del servidor
APP_HOST=
APP_PORT=
APP_URL="http://${APP_HOST}:${APP_PORT}"

# Configuración de la base de datos (CockroachDB o PostgreSQL)
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=

# Configuración de Redis para caché
CACHE_ADDRESS=
CACHE_PASSWORD=

# API para calificaciones de acciones
STOCK_API_URL=
STOCK_API_KEY=

# API para datos financieros
FINANCIAL_API_URL=
FINANCIAL_API_KEY=

# API para Gemini
GEMINI_API_KEY=
GEMINI_API_URL=
```

## Instalación y Ejecución

### Requisitos Previos

- Go 1.19 o superior
- PostgreSQL o CockroachDB
- Redis para caché
- Claves de API para los servicios integrados

### Instalación

1. Clona el repositorio
2. Configura las variables de entorno en un archivo `.env`
3. Instala las dependencias:

```bash
go mod download
```

### Ejecución

Para iniciar el servidor:

```bash
go run main.go
```

El servidor estará disponible en http://localhost:3001 (o el puerto configurado).

## Comandos CLI

La aplicación incluye algunos comandos CLI que se pueden ejecutar para tareas administrativas o de mantenimiento:

```bash
go run main.go [comando]
```

Comandos disponibles:

- `load-stock`: Carga la información base (tickers, brokers y recomendaciones) con la cual se sustenta el proyecto. Es necesario ejecutar este comando antes de utilizar la API por primera vez para tener datos iniciales.

Ejemplo:

```bash
go run main.go load-stock
```

## Licencia

Este proyecto está licenciado bajo los términos de la [Licencia MIT](https://opensource.org/licenses/MIT).
