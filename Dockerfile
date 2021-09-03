FROM golang:1.16

COPY . /bot
WORKDIR /bot

RUN go mod download
RUN go build -o bot ./cmd/main.go

CMD ["./bot"]
