FROM golang:1.19.5-alpine

LABEL maintainer="manujurado1@correo.ugr.es" \
      version="v0.0.5"

RUN apk add build-base

WORKDIR /app/test

COPY go.mod ./

RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT [ "task", "test" ]