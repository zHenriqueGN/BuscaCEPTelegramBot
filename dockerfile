FROM golang:1.21.2-alpine3.18

WORKDIR /bot

COPY . .

RUN go mod download

RUN go build -o bot cmd/bot/main.go

ENTRYPOINT [ "./bot" ]
