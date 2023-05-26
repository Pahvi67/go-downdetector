FROM golang:1.20.4-alpine3.18 AS builder
RUN mkdir /build
WORKDIR /build
COPY go.mod go.sum index.go ./
RUN go mod download
RUN go build -o main .

FROM alpine
COPY --from=builder /build/main /app/
COPY start.sh /app/
WORKDIR /app
CMD ["./start.sh"]
