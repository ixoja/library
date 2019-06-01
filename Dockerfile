FROM golang

ENV GOBIN $GOPATH/bin

ADD . /go/src/github.com/ixoja/library
RUN go install /go/src/github.com/ixoja/library/internal/cmd/library-server
WORKDIR /go/src/github.com/ixoja/library
ENTRYPOINT /go/bin/library-server --port 8090 --host 0.0.0.0

# serving HTTP of 8090
EXPOSE 8090