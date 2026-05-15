FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY services/posts/ services/posts/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o posts ./services/posts

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/posts .
COPY services/posts/etc/ etc/
EXPOSE 8882
CMD ["./posts"]
