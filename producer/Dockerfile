FROM golang:1.15.10-alpine3.13
WORKDIR $GOPATH/src/mypackage/producer
COPY [ "producer.go", "../go.mod", "./"]
RUN go get -d -v ./

RUN go build -o /go/bin/producer

ENTRYPOINT ["/go/bin/producer"]
