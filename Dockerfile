# syntax=docker/dockerfile:1
FROM golang:alpine

RUN apk add --no-cache musl-dev gcc 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go get github.com/mattn/go-sqlite3
RUN CGO_ENABLED=1 GOOS=linux go build -o /url-shortener
EXPOSE 8080
CMD ["/url-shortener"]
