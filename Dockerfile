FROM alpine:3.5

ADD . /go-shaft
RUN \
  apk add --update git go make gcc musl-dev linux-headers && \
  (cd go-shaft && make geth)                           && \
  cp go-shaft/build/bin/geth /usr/local/bin/           && \
  apk del git go make gcc musl-dev linux-headers          && \
  rm -rf /go-shaft && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303
EXPOSE 30303/udp

ENTRYPOINT ["geth"]
