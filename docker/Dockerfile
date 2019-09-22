FROM alpine:3.6

RUN apk add --no-cache ca-certificates

ADD pulsar-operator /usr/local/bin/pulsar-operator

RUN adduser -D pulsar-operator
USER pulsar-operator
