FROM golang:1.19.5-alpine

LABEL maintainer="manujurado1@correo.ugr.es" \
      version="v0.0.5"

WORKDIR /app/test

RUN adduser --disabled-password --uid 1001 test && chown test:test /app/test

COPY go.mod go.sum Taskfile.yml ./

COPY pkg ./pkg

RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT [ "task", "test" ]