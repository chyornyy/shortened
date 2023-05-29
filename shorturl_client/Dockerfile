FROM golang:1.20-alpine
RUN apk add --no-cache inotify-tools
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./shorturl_client/main.go
CMD ["./main"]
EXPOSE 8080
ENV GIN_MODE=debug