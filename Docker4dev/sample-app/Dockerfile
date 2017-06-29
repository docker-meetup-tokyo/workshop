# Base image
FROM golang:1.8

ENV HTTP_PROXY="proxy.rdh.ecl.ntt.co.jp:20066"
ENV HTTPS_PROXY="proxy.rdh.ecl.ntt.co.jp:20066"

# Fetch dependencies
RUN go get github.com/buger/goterm

# Copy files from context
COPY . /vcoin

# Declare volume, so can be mounted from outside.
VOLUME /vcoin

# Declare it as workdir, On container creation, by default this will be current dir.
WORKDIR /vcoin




