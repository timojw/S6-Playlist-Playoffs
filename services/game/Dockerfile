# Stage 1: Build the Go application
FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy all files to the container
COPY . .

# Change to the service directory
WORKDIR /app/services/game

# Download dependencies
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/game-service ./cmd

# Stage 2: Create the runtime image
FROM scratch
COPY --from=builder /app/game-service /game-service

EXPOSE 8081
CMD ["/game-service"]
