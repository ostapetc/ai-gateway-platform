FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY services/comments/ services/comments/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o comments ./services/comments

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/comments .
COPY services/comments/etc/ etc/
EXPOSE 8883 8884
CMD ["./comments"]
