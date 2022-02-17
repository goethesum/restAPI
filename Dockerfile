FROM golang:1.17 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=LINUX go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]
