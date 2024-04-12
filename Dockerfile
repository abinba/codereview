FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add curl ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/start.sh .
COPY --from=builder /app/migrations ./migrations
RUN chmod +x ./start.sh

EXPOSE 8080

CMD ["/bin/sh", "./start.sh"]
