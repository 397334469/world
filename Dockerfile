FROM golang:latest as build-env

ENV GO111MODULE=on
ENV BUILDPATH=github.com/kplcloud/world
#ENV GOPROXY=goproxy.io
ENV GOPATH=/go
RUN mkdir -p /go/src/${BUILDPATH}
COPY ./ /go/src/${BUILDPATH}
RUN cd /go/src/${BUILDPATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v

FROM alpine:latest

COPY --from=build-env /go/bin/world /go/bin/world
WORKDIR /go/bin/
CMD ["/go/bin/world"]
