FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy the voting service source code into the build context
COPY . .

RUN cd services/voting && CGO_ENABLED=0 GOOS=linux go build -o voting-service ./cmd

FROM alpine:3.19.1
WORKDIR /app
COPY --from=builder /app/services/voting/voting-service /app/voting-service

EXPOSE 8082
CMD ["./voting-service"]
