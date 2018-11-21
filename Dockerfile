FROM golang:1.11.2

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ENV maillib /go/src/github.com/JobTool/mail-thread-reader
RUN mkdir -p $maillib
WORKDIR $maillib
ADD . Gopkg.toml
ADD . Gopkg.lock
RUN dep ensure
