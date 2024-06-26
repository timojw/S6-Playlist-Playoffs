FROM golang:1.22.2-alpine AS builder
WORKDIR /app
RUN mkdir -p services/game
RUN mkdir data

COPY . .
RUN cd services/game && CGO_ENABLED=0 GOOS=linux go build -o game-service ./cmd

# Build
FROM alpine:3.19.1
WORKDIR /app
COPY --from=builder /app/ /app/
RUN mv /app/services/game/game-service /app/services

EXPOSE 8081
CMD ["services/game-service"]
