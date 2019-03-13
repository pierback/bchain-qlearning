FROM golang:1.9.1
LABEL maintainer="fabian.pieringer@gmail.com"

COPY . /go/src/github.com/pierback/bchain-qlearning

WORKDIR /go/src/github.com/pierback/bchain-qlearning

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["cmd"]
