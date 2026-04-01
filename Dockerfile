# Build Stage
FROM golang:1.26-alpine AS builder
WORKDIR /app

ENV GOPROXY=https://proxy.golang.org,direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o boilerplate-service ./cmd

# Run Stage
FROM alpine:latest
WORKDIR /root/
ENV config=docker
COPY --from=builder /app/boilerplate-service .
COPY --from=builder /app/config ./config
COPY --from=builder /app/docs ./docs
# Expose the port Fiber listens on internally
EXPOSE 3000 
CMD ["./boilerplate-service"]