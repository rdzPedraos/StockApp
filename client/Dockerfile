FROM node:20-alpine AS builder

WORKDIR /app

# Instalar pnpm
RUN corepack enable && corepack prepare pnpm@10.8.1 --activate

# Copiar archivos de dependencias
COPY package.json pnpm-lock.yaml ./

# Instalar dependencias
RUN pnpm install --frozen-lockfile

# Copiar el resto del código
COPY . .

# Compilar la aplicación
RUN pnpm build

# Imagen final
FROM node:20-alpine

WORKDIR /app

# Instalar pnpm
RUN corepack enable && corepack prepare pnpm@10.8.1 --activate

# Copiar los archivos necesarios de la etapa de compilación
COPY --from=builder /app/.output ./.output
COPY --from=builder /app/package.json ./package.json

# Exponer el puerto
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["node", ".output/server/index.mjs"] 