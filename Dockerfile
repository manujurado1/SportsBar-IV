FROM golang:1.20rc2-alpine3.17

LABEL maintainer="manujurado1@correo.gur.es"

WORKDIR /app/test

RUN adduser --disabled-password --uid 1001 test && chown test:test /app/test

COPY . .

RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT [ "task", "test" ]