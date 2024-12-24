FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o myapp .
FROM alpine:latest
COPY --from=builder /app/myapp .
COPY config.json .
EXPOSE 8080
CMD ["./myapp"]