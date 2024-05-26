FROM golang:1.22.2-alpine AS builder

WORKDIR /app

RUN mkdir -p services/game   
RUN mkdir data

COPY services/game services/game
# Build
RUN cd services/game && CGO_ENABLED=0 GOOS=linux go build -o game-service ./cmd

# Use a minimal base image for the final image
FROM alpine:3.19.1
WORKDIR /app
# Copy the built binary from the builder stage
COPY --from=builder /app/ /app/
RUN mv /app/services/game/game-service /app/services

# Expose the port
EXPOSE 6000

# Run the binary
CMD ["services/game-service"]