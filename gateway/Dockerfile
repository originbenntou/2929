FROM golang:1.14.0 as builder

ENV WORK_DIR /go/src/github.com/originbenntou/2929BE
ENV MAIN_PATH $WORK_DIR/gateway

WORKDIR $WORK_DIR

ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $WORK_DIR

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $MAIN_PATH

FROM alpine:latest
ENV TZ Asia/Tokyo
RUN mkdir -p /var/log/ap && touch /var/log/ap/ap.log
COPY --from=builder /go/bin/gateway /go/bin/gateway
ENTRYPOINT ["/go/bin/gateway"]
