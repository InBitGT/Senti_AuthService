# Etapa 1: build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app main.go

# Etapa 2: imagen final ligera
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8081
CMD ["./app"]
