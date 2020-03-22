FROM golang:latest AS cached_deps
LABEL author='Daniel Lee'

WORKDIR /go/app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -v -o /go/bin/telegrambot ./cmd/telegrambot/main.go

FROM cached_deps
CMD /go/bin/telegrambot