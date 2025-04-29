FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy dependencies and download them
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Compile the application
# CGO_ENABLED=0 is used to build a static binary, without dependencies on C libraries, better for alpine
# GOOS=linux is used to build for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final image
FROM alpine:3.19

WORKDIR /app

# Install necessary dependencies
RUN apk --no-cache add \
    # Certificates for SSL/TLS
    ca-certificates \ 
    # Timezone data
    tzdata

# Copy the compiled executable
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Command to run the application
CMD ["./main"] 