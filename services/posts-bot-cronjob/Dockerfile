FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY services/users/grpc/ services/users/grpc/
COPY services/posts/grpc/ services/posts/grpc/
COPY services/posts-bot-cronjob/ services/posts-bot-cronjob/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o posts-bot-cronjob ./services/posts-bot-cronjob

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/posts-bot-cronjob .
COPY services/posts-bot-cronjob/etc/ etc/
ENTRYPOINT ["./posts-bot-cronjob"]
