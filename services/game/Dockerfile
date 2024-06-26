FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy only the relevant part of the repository
COPY services/game /app/services/game

# Build the Go application
RUN cd /app/services/game && CGO_ENABLED=0 GOOS=linux go build -o game-service ./cmd

FROM alpine:3.19.1
WORKDIR /app
COPY --from=builder /app/services/game/game-service /app/game-service

EXPOSE 8081
CMD ["./game-service"]
