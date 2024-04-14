FROM golang:1.22.2-alpine

WORKDIR /game

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o game main.go

ENTRYPOINT ["./game"]