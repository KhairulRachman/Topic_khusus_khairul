# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /app


# Copy seluruh project
COPY . .
RUN go mod init payment-options
RUN go mod tidy
RUN go mod vendor
RUN go mod verify

# Build binary dari cmd/main.go
RUN go build -o payment-options cmd/main.go

# Copy docker-compose.prod.yml to the image
COPY docker-compose.prod.yml /app/docker-compose.prod.yml

# Run stage
FROM alpine:latest

# Install SSL certificates (kalau kamu pakai HTTPS/Go HTTP client)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary dari stage sebelumnya
COPY --from=build /app/payment-options .
# Copy docker-compose.prod.yml from the build stage
COPY --from=build /app/docker-compose.prod.yml /app/docker-compose.prod.yml


# Jalankan binary
CMD ["./payment-options"]
