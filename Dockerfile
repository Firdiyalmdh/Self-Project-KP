FROM golang:1.18 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=LINUX go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 2500

CMD ["./app"]