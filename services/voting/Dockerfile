FROM golang:1.22.2-alpine AS builder
WORKDIR /app
RUN mkdir -p services/voting
RUN mkdir data

COPY services/voting /app/services/voting
RUN cd services/voting && CGO_ENABLED=0 GOOS=linux go build -o voting-service ./cmd

# Build
FROM alpine:3.19.1
WORKDIR /app
COPY --from=builder /app/ /app/
RUN mv /app/services/voting/voting-service /app/services/voting

EXPOSE 8082
CMD ["services/voting/voting-service"]
