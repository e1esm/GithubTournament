FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add git && apk add bash

COPY . .

RUN go mod download && go mod tidy

ENV GOBIN /go/bin

RUN go build -o main ./cmd/bot/main.go

CMD ["./main"]