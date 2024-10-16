# Build stage
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

# Run stage
# FROM alpine:latest AS runner

FROM golang:1.23.1-alpine AS runner

WORKDIR /app

COPY --from=builder /app/main .

COPY local /home/longln/SourceCode/github.com/longln/go-simplebank

RUN chmod +x ./main

EXPOSE 8080

# RUN ls /home/longln/SourceCode/github.com/longln/go-simplebank/local

# CMD ["/app/main"]