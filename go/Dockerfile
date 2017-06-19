FROM golang
WORKDIR /app
ENV SRC_DIR=/go/src/github.com/kainobor/spherq/go/
ADD . $SRC_DIR
RUN cd $SRC_DIR; go get; go build -o spherq; cp spherq /app/
ENTRYPOINT ["./spherq"]