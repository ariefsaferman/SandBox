#build stage
FROM golang:1.18-alpine AS builder
LABEL maintainer="Digitalent"
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . ./ 
RUN go mod download 
ENV GIN_MODE=release
RUN go build -o /library-api

EXPOSE 8080

CMD ["/library-api"]

