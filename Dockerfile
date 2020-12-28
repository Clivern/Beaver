FROM golang:1.15.6 as builder

ENV GO111MODULE=on

ARG BEAVER_VERSION=1.2.2

WORKDIR $GOPATH/src/github.com/clivern/beaver

ADD . .

RUN git checkout tags/$BEAVER_VERSION

RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o /go/bin/beaver .

RUN mkdir -p /go/logs/beaver && mkdir -p /go/configs/beaver

# Build a small image
FROM alpine:3.12
RUN  apk --no-cache add ca-certificates

# Copy our static executable
COPY --from=builder /go/bin/beaver /go/bin/beaver
COPY --from=builder /go/logs/beaver /go/logs/beaver
COPY --from=builder /go/configs/beaver /go/configs/beaver

WORKDIR /go/bin/

EXPOSE 8080

CMD ["./beaver"]