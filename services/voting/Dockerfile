# Stage 1: Build the Go application
FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files
COPY services/voting/go.mod services/voting/go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY services/voting .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/voting-service ./cmd

# Stage 2: Create the runtime image
FROM scratch
COPY --from=builder /app/voting-service /voting-service

EXPOSE 8082
CMD ["/voting-service"]
