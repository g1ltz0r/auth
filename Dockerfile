FROM golang:1.23.2-alpine3.20 AS builder

COPY . /github.com/g1ltz0r/auth/grpc/source
WORKDIR /github.com/g1ltz0r/auth/grpc/source

RUN go mod download
RUN go build -o ./bin/auth_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/g1ltz0r/auth/grpc/source/bin/auth_server .

CMD ["./auth_server"]