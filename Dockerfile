FROM golang:latest
RUN go install github.com/cosmtrek/air@latest

WORKDIR /go/src/app

CMD ["air", "-c", ".air.toml"]
