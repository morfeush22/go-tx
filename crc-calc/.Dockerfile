FROM golang:latest

WORKDIR /go/src/github.com/morfeush22/go-tx/crc-calc
COPY . .
RUN go get -d -v ./...
RUN rm -r /go/src/github.com/morfeush22/go-tx/crc-calc
EXPOSE 8080
