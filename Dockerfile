FROM golang:alpine

RUN apk add -U tzdata
ENV GO111MODULE=on
ENV APP_HOME /usr/src/app
ENV TZ=Asia/Jakarta
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

ADD . "$APP_HOME"
WORKDIR "$APP_HOME"

RUN go mod tidy
RUN go build app/main.go

EXPOSE 8090

#CMD ["make", "start"]
ENTRYPOINT [ "./main" ]
