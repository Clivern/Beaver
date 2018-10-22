FROM golang:1.11.1

RUN curl https://raw.githubusercontent.com/golang/dep/v0.5.0/install.sh | sh

RUN mkdir -p /go/src/github.com/clivern/beaver/

ADD . /go/src/github.com/clivern/beaver/

WORKDIR /go/src/github.com/clivern/beaver

RUN dep ensure

RUN go build -o beaver beaver.go

EXPOSE 8080

CMD ["./beaver"]