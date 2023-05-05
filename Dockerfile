FROM golang:1.20.3-alpine3.17

WORKDIR /go/src/github.com/AkankshaNichrelay/TableFinder

ENV GO111MODULE=on
COPY ./ ./
RUN go mod download
RUN go build -o ./bin/tablefinder ./cmd/tablefinder

CMD ["./bin/tablefinder"]