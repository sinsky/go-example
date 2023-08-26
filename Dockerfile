FROM golang:1.21 AS builder

ENV CGO_ENABLED=0 \
  GOOS=linux

WORKDIR /app

COPY main.go go.mod ./

RUN go build -o easy-server

FROM alpine:latest

RUN addgroup -S gouser && adduser -S gouser -G gouser
USER gouser:gouser

COPY --from=builder --chown=gouser:gouser /app/easy-server /easy-server

EXPOSE 8080

CMD ["/easy-server"]
