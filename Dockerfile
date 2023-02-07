# syntax=docker/dockerfile:1

FROM golang:1.18.1-alpine3.14

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /pharmacy-server

CMD ["/pharmacy-server"]
