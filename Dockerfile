FROM golang:1.9.1
LABEL maintainer="fabian.pieringer@gmail.com" 

COPY . /go/src/bchain-qlearning
# COPY ./internal /go/src/bchain-qlearning/internal
# COPY ./internal/learning /go/src/bchain-qlearning/internal/learning

WORKDIR /go/src/bchain-qlearning
COPY ./cmd .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["bchain-qlearning"]
