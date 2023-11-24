FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main
  # Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 3007
CMD [ "/app/main" ]