# Use the official Go image as build environment
FROM golang:1.24-alpine AS builder

# Set working directory inside container
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./userapilab01

FROM gcr.op/distroless/base-debian12

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/userapilab01 ./main
COPY --from=builder /app/.env .

CMD ["./main"]
