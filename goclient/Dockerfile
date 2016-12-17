FROM golang:latest
MAINTAINER Julien Fabre <ju.pryz@gmail.com>

RUN git clone https://github.com/edenhill/librdkafka.git /librdkafka

WORKDIR /librdkafka

RUN ./configure --prefix /usr
RUN make
RUN make install

RUN go get -d github.com/confluentinc/confluent-kafka-go || true
WORKDIR $GOPATH/src/github.com/confluentinc/confluent-kafka-go/kafka
RUN go install

