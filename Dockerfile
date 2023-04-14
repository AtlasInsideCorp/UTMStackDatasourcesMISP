FROM golang:alpine

RUN apk update 

WORKDIR /usr/src/app

COPY config /usr/src/app/config
COPY controller /usr/src/app/controller
COPY models /usr/src/app/models
COPY serv /usr/src/app/serv
COPY utils /usr/src/app/utils
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod tidy
ENTRYPOINT go run .
