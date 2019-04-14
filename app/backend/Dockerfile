FROM golang:latest

WORKDIR /go/Backend

ADD . .

WORKDIR /go/Backend/src

RUN go build -o backendServer

ENTRYPOINT ./backendServer

EXPOSE 8080
