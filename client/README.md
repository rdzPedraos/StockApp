# Stock App - Cliente

## Descripción

Stock App es una aplicación web moderna desarrollada con Vue 3 y Nuxt 3 que proporciona herramientas para análisis y toma de decisiones financieras. La aplicación permite a los usuarios consultar información detallada sobre acciones bursátiles, obtener recomendaciones de inversión y visualizar datos de mercado para facilitar la toma de decisiones financieras informadas.

## Tecnologías Principales

-   **Framework**: [Nuxt 3](https://nuxt.com/)
-   **UI Framework**: [Vue 3](https://vuejs.org/)
-   **Gestión de Estado**: [Pinia](https://pinia.vuejs.org/)
-   **Estilos**: [Tailwind CSS](https://tailwindcss.com/) + [Nuxt UI](https://ui.nuxt.com/)
-   **Peticiones HTTP**: Axios
-   **Lenguaje**: TypeScript

## Funcionalidades Principales

-   Visualización de información actualizada de acciones (tickers)
-   Análisis detallado de empresas
-   Recomendaciones de compra/venta de acciones
-   Gráficos interactivos para análisis técnico
-   Recomendaciones basadas en IA

## Estructura del Proyecto

El proyecto sigue la estructura estándar de Nuxt 3:

```
├── assets/                # Recursos estáticos (CSS, imágenes)
├── components/            # Componentes Vue reutilizables
├── composables/           # Composables de Vue para lógica reutilizable
├── layouts/               # Layouts para las páginas
├── pages/                 # Páginas de la aplicación
├── plugins/               # Plugins de Nuxt
├── public/                # Archivos públicos
├── server/                # Middleware y API de servidor
├── shared/                # Tipos y utilidades compartidas
├── stores/                # Stores de Pinia para gestión de estado
├── utils/                 # Utilidades y funciones auxiliares
├── .env                   # Variables de entorno
├── .env.example           # Ejemplo de variables de entorno
├── app.config.ts          # Configuración de la aplicación
├── app.vue                # Componente principal
├── nuxt.config.ts         # Configuración de Nuxt
└── package.json           # Dependencias y scripts
```

## Páginas Principales

-   **Home**: Muestra un resumen de los tickers y sus precios actuales
-   **Ticker**: Muestra información detallada de un ticker específico, incluidos gráficos, recomendaciones y datos de la empresa

## Integración con Backend

La aplicación se comunica con una API backend desarrollada en Go que proporciona los datos necesarios. La documentación completa de la API se encuentra en el directorio `/api`.

Características principales de la API:

-   Obtención de información detallada de tickers
-   Recomendaciones de inversión
-   Análisis de mercado con IA
-   Datos históricos de precios

## Configuración

### Variables de Entorno

Copia el archivo `.env.example` a `.env` y configura las siguientes variables:

```
API_URL=http://localhost:3001
```

### Instalación

```bash
# Instalar dependencias con pnpm (recomendado)
pnpm install

# Alternativas
npm install
yarn install
```

### Desarrollo

Inicia el servidor de desarrollo:

```bash
# Con pnpm
pnpm dev

# Alternativas
npm run dev
yarn dev
```

La aplicación estará disponible en http://localhost:3000.

### Producción

Construye la aplicación para producción:

```bash
# Con pnpm
pnpm build

# Alternativas
npm run build
yarn build
```

Previsualiza la versión de producción:

```bash
# Con pnpm
pnpm preview

# Alternativas
npm run preview
yarn preview
```

## Componentes Principales

-   **ChartStock**: Visualiza gráficos interactivos para análisis técnico
-   **StockInfoCard**: Muestra información resumida de una acción
-   **RecommendationCard**: Presenta recomendaciones de inversión
-   **TickerInfo**: Muestra información detallada de un ticker

## Gestión de Estado con Pinia

La aplicación utiliza Pinia para la gestión del estado. Los principales stores incluyen:

-   **tickerStore**: Gestiona los datos de los tickers, incluyendo información de la empresa y recomendaciones

## Licencia

Este proyecto está licenciado bajo los términos de la Licencia MIT.
