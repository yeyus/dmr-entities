FROM golang:1.12-alpine as builder

RUN apk update && \
    apk upgrade && \
    apk add protobuf && \
    apk add git && \
    apk add make

RUN go get google.golang.org/grpc && \
    go get github.com/golang/protobuf/protoc-gen-go

COPY . /app

RUN cd /app && \
    go get && \
    make all

# ---
FROM scratch as syncjob

COPY  --from=builder /app/bin/syncjob /app/syncjob
CMD ["./app/syncjob"]

# ---
FROM scratch as server

COPY  --from=builder /app/bin/server /app/server
CMD ["./app/server"]