FROM golang:alpine as Client
RUN apk update && apk add --no-cache git inotify-tools jq

COPY . /src

RUN cd /src/Client && go build -o /go/Client client.go
RUN mv /src/Client/docker-entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]