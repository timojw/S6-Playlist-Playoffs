FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy the game service source code
COPY services/game /app/services/game

# Build the Go application
RUN cd services/game && CGO_ENABLED=0 GOOS=linux go build -o game-service ./cmd

# Use a minimal image for the final build
FROM alpine:3.19.1
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/services/game/game-service /app/game-service

# Expose the application's port
EXPOSE 8081

# Command to run the application
CMD ["./game-service"]
