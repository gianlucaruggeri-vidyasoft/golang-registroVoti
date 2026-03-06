FROM golang:1.26-alpine AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -trimpath -ldflags="-s -w" -o app ./src/cmd/api
FROM alpine:3.21 AS final
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app .
EXPOSE 3000
CMD ["./app"]