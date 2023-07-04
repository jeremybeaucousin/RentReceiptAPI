FROM golang:1.17.13-alpine3.16

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN apk update && apk add git
RUN go build -o main .

CMD ["/app/main"]