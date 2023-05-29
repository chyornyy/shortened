FROM golang:1.20-alpine
RUN apk add --no-cache inotify-tools
RUN apk add --no-cache git
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./shorturl_server/main.go
RUN apk add --no-cache postgresql-client
ENV DATABASE=postgresql
CMD ["./main"]
EXPOSE 50051