FROM golang

RUN mkdir -p /go/src/github.com/klim0v/go-kit-example

ADD . /go/src/github.com/klim0v/go-kit-example

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/klim0v/go-kit-example/adder/cmd  -watch github.com/klim0v/go-kit-example/adder
