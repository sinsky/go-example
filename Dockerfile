FROM golang:1.21 AS builder

WORKDIR /app

COPY main.go go.mod ./

RUN CGO_ENABLED=0 GOOS=linux go build -o easy-server

FROM alpine:latest

COPY --from=builder /app/easy-server /easy-server

USER gouser:gouser

EXPOSE 8080

CMD ["/easy-server"]
