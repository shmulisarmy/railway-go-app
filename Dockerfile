FROM golang:1.20 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root
COPY --from=builder /app/app .

CMD ["./app"]
