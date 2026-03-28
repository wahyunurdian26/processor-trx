# Dockerfile for trx-processor
FROM golang:1.24-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
COPY util/ ./util/
COPY cp-proto/ ./cp-proto/
COPY trx-processor/ ./trx-processor/

RUN go mod download

RUN go build -o main trx-processor/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
