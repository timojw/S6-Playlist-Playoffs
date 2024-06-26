FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy only the relevant part of the repository
COPY services/voting .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o voting-service ./cmd

FROM scratch

COPY --from=builder /app/voting-service /voting-service

EXPOSE 8081
CMD ["/voting-service"]
