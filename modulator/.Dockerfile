FROM golang:latest

WORKDIR /go/src/github.com/morfeush22/go-tx/modulator
COPY . .
RUN go get -d -v ./...
RUN rm -r /go/src/github.com/morfeush22/go-tx/modulator
EXPOSE 8080
