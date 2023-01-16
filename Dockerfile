FROM golang:1.19.5-alpine

LABEL maintainer="manujurado1@correo.ugr.es" \
      version="v0.0.5"

WORKDIR /app

RUN apk add build-base && adduser -D -u 1001 test && chown test /app

USER test

WORKDIR /app/test

COPY go.mod ./

RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT [ "task", "test" ]