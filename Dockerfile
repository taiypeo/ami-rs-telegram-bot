FROM golang

COPY . /go/src/github.com/taiypeo/ami-rs-telegram-bot
WORKDIR /go/src/github.com/taiypeo/ami-rs-telegram-bot
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/ami-rs-telegram-bot
