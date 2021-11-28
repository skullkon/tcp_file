FROM golang:1.17.2-alpine3.14 as builder

COPY go.mod go.sum /go/src/tcp_file/

WORKDIR /go/src/tcp_file/

COPY . /go/src/tcp_file/

RUN CGO_ENABLED=0 GOOS=linux go build -o build/tcp_file tcp_file/cmd/server

FROM alpine
#RUN apk update && apk add —no-cache git ca-certificates tzdata && update-ca-certificates
#
COPY —from=builder ./go/src/tcp_file/build/tcp_file /usr/bin/tcp_file

EXPOSE 8080 8081

ENTRYPOINT ["/usr/bin/tcp_file"]