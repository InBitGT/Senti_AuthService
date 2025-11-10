# Etapa 1: build
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/main.go

# Etapa 2: imagen final ligera
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8000
CMD ["./app"]
