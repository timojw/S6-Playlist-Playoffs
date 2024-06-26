FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy only the relevant part of the repository
COPY services/game .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o game-service ./cmd

FROM scratch

COPY --from=builder /app/game-service /game-service

EXPOSE 8081
CMD ["/game-service"]
