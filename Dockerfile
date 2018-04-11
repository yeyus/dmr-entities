FROM golang:1.8

WORKDIR /app

COPY bin/dmr-entities-linux-amd64 dmr-entities

CMD ["./dmr-entities"]