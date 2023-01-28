# Golang Base Image

FROM golang:1.18-alpine

# Add Maintainer Info

LABEL maintainer="Digitalent"

# Install git.

# Git is required for fetching the dependencies.

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup Folder

WORKDIR /app

# Copy source from current directory to the working Directory inside the container

COPY go.mod ./

COPY go.sum ./

COPY ./db/db-dump.sql ./docker-entrypoint-initdb.d/db-dump.sql

# Ensure each dependecies is valid

RUN go mod download

# Copy everything else

COPY . ./

ENV GIN_MODE=release

RUN go build -o ./server ./server.go

EXPOSE 8080

CMD [ "./server" ]