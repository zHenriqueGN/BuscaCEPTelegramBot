FROM golang:1.21.2-alpine3.18 AS build

WORKDIR /bot
COPY . .
RUN go mod download
RUN go build -o bot cmd/bot/main.go


FROM alpine:3.18 AS binary

COPY --from=build /bot/bot .
ENTRYPOINT [ "./bot" ]
