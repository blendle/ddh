FROM google/golang
MAINTAINER Koen Bollen <koen@blendle.com>

RUN go get github.com/go-martini/martini
RUN go get github.com/fsouza/go-dockerclient

ENV DOCKERSOCKET unix:///var/run/docker.sock
ENV ENDPOINT /
ENV IMAGE ubuntu
ENV CONTAINER_NAME ubuntu-test
ENV CMD sleep 500

#ENV USERNAME you
#ENV PASSWORD secret

ENV PASS_ENV PASS_ENV


WORKDIR /gopath/src/app
ADD . /gopath/src/app/

ENTRYPOINT go run /gopath/src/app/main.go
