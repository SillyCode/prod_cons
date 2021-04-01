FROM golang:1.15.10-alpine3.13
WORKDIR $GOPATH/src/mypackage/consumer
COPY [ "consumer.go", "../go.mod", "./"]
RUN go get -d -v ./

RUN go build -o /go/bin/consumer

ENTRYPOINT ["/go/bin/consumer"]
