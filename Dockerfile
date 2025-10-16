# Multi-stage build for minimal runtime image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Pre-download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server ./cmd/main.go

FROM alpine:3.19 AS final
WORKDIR /app

# Copy binary
COPY --from=builder /app/server /app/server

# Expose service port
EXPOSE 8080

# Use release mode for gin
ENV GIN_MODE=release

# Run the server
ENTRYPOINT ["/app/server"]