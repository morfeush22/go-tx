FROM golang:latest

WORKDIR /go/src/crc-calc
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["crc-calc"]
