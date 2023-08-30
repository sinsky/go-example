FROM golang:1.21 AS builder

ENV CGO_ENABLED=0 \
  GOOS=linux

WORKDIR /app

COPY . ./

RUN go mod download

RUN go vet -v

RUN go build -o easy-server

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/easy-server /easy-server

EXPOSE 8080

CMD ["/easy-server"]
