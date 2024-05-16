FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add curl ca-certificates

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/migrations /app/migrations
RUN chmod +x start.sh
RUN dos2unix start.sh

EXPOSE 8080

CMD ["/bin/sh", "/app/start.sh"]
