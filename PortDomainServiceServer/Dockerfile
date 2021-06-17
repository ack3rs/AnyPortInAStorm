FROM golang:alpine as portdomainserviceserver
RUN apk update && apk add --no-cache git inotify-tools jq

COPY . /src

# RUN cd /src && GIT_COMMIT=$(git rev-list -1 HEAD) && VERSION=$(git describe --tags --abbrev=0) && go build server.go -o /go/Server -ldflags "-X main.GitCommit=$GIT_COMMIT" -ldflags "-X main.VERSION=$VERSION"
RUN cd /src/PortDomainServiceServer && go build -o /go/PortDomainServiceServer Server.go
RUN mv /src/PortDomainServiceServer/docker-entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]