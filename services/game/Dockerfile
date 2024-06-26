# Stage 1: Build the Go application
FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files
COPY services/game/go.mod services/game/go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o game-service ./cmd

# Stage 2: Create the final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/game-service .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./game-service"]
