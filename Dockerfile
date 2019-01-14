FROM golang:1.11.1

ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/clivern/beaver/

ADD . /go/src/github.com/clivern/beaver/

WORKDIR /go/src/github.com/clivern/beaver

RUN go build -o beaver beaver.go

EXPOSE 8080

CMD ["./beaver"]