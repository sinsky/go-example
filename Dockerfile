FROM golang:1.21 AS builder

WORKDIR /app

COPY main.go go.mod go.sum ./

RUN go build -o easy-server

FROM scratch

COPY --from=builder /app/easy-server /

EXPOSE 8080

CMD ["/easy-server"]
