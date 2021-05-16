FROM golang:alpine

WORKDIR /main
COPY . /main

RUN go mod download
RUN go build -o unischeme

ENTRYPOINT ["/main/unischeme"]
