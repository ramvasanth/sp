FROM golang:latest

WORKDIR /go/src/github.com/ramvasanth/sp
RUN apt-get update
RUN apt-get install vim -y
COPY . .