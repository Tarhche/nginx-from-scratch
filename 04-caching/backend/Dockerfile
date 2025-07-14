# building stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o backend main.go

# runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend ./backend

EXPOSE 8080

CMD ["/app/backend"]
