FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy the voting service source code
COPY services/voting /app/services/voting

# Build the Go application
RUN cd services/voting && CGO_ENABLED=0 GOOS=linux go build -o voting-service ./cmd

# Use a minimal image for the final build
FROM alpine:3.19.1
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/services/voting/voting-service /app/voting-service

# Expose the application's port
EXPOSE 8082

# Command to run the application
CMD ["./voting-service"]
