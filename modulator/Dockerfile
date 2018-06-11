FROM golang:latest

WORKDIR /go/src/github.com/morfeush22/go-tx/modulator
COPY . .
RUN go get -d -v ./...
RUN go build -v .
EXPOSE 8080
CMD ["./modulator"]
