FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git

WORKDIR /build/server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server shorturl_server/shorturl_server.go

WORKDIR /build/client
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o client shorturl_client/shorturl_client.go


FROM alpine:3.13

RUN apk add --no-cache postgresql-client

WORKDIR /app
COPY --from=builder /build/server/server .
COPY --from=builder /build/client/client .
ENV GIN_MODE=release
ENV DATABASE="in-memory"
EXPOSE 50051 8080
CMD ./server -database=$DATABASE & ./client