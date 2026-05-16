FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY services/comments/ services/comments/
COPY services/posts/ services/posts/
COPY services/comments-bot-cronjob/ services/comments-bot-cronjob/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o comments-bot-cronjob ./services/comments-bot-cronjob

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/comments-bot-cronjob .
COPY services/comments-bot-cronjob/etc/ etc/
ENTRYPOINT ["./comments-bot-cronjob"]
