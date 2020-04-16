FROM golang:1.14.0

ENV WORK_DIR /go/src/github.com/originbenntou/2929BE
ENV MAIN_PATH $WORK_DIR/gateway

WORKDIR $WORK_DIR

ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $WORK_DIR

# ホットリロード
WORKDIR $MAIN_PATH
RUN go get github.com/pilu/fresh
CMD ["fresh", "-c", "docker-local/fresh/runner.conf"]
