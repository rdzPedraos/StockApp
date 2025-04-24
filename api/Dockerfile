FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiar archivos de dependencias y descargarlas
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Imagen final
FROM alpine:3.19

WORKDIR /app

# Instalar dependencias necesarias
RUN apk --no-cache add ca-certificates tzdata

# Copiar el ejecutable compilado
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"] 