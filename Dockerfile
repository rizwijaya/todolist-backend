FROM golang:alpine

ENV APP_HOME /usr/src/app

ADD . "$APP_HOME"
WORKDIR "$APP_HOME"

RUN go mod tidy
RUN go build app/main.go

EXPOSE 3030

ENTRYPOINT [ "./main" ]