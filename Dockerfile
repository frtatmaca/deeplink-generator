FROM golang:1.17

RUN mkdir /deeplink

ADD . /deeplink

WORKDIR /deeplink

VOLUME /logs

RUN go build -o server .

CMD ["/deeplink/server"]