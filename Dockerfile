FROM golang:1.12

RUN mkdir -p /go/src/github.com/ramvasanth/sp
ADD . /go/src/github.com/ramvasanth/sp
WORKDIR /go/src/github.com/ramvasanth/sp
RUN chmod +x ./script/build
RUN ./script/build
CMD /go/src/github.com/ramvasanth/sp/bin/worker