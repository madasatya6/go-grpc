FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./app/api/main.go

FROM alpine:3.14

RUN apk update && apk --no-cache add curl ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app/api /app/api
COPY --from=builder /app/asset /app/asset

EXPOSE 8080
ENTRYPOINT ["/app/api"]
