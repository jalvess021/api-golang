FROM golang:1.23.3

WORKDIR /go/app

RUN apt-get update && apt-get install -y librdkafka-dev

COPY . .

RUN go mod tidy

CMD ["tail", "-f", "/dev/null"]