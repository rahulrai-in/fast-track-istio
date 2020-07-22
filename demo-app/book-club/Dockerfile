FROM golang:1.13-alpine

LABEL maintainer="Rahul Rai <rahul@rahul-rai.com> & Tarun Pabbi <tarun.pabbi@outlook.com>"

WORKDIR /go/src

COPY . /go/src

RUN cd /go/src/cmd/web/ && go build -o main

EXPOSE 8080

ENTRYPOINT ["/go/src/cmd/web/main", "-addr=:8080"]
