FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY services/users/ services/users/
COPY services/users-bot-cronjob/ services/users-bot-cronjob/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o users-bot-cronjob ./services/users-bot-cronjob

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/users-bot-cronjob .
COPY services/users-bot-cronjob/etc/ etc/
ENTRYPOINT ["./users-bot-cronjob"]
