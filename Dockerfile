FROM golang:1.17

RUN mkdir -p $GOPATH/src/github.com/note_keeper
ADD . $GOPATH/src/github.com/note_keeper

WORKDIR $GOPATH/src/github.com/note_keeper

RUN go get -u github.com/lib/pq@v1.10.2
RUN go mod tidy

EXPOSE 8080