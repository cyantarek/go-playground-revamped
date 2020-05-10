FROM golang:latest as builder

ARG BIN_NAME=playground

COPY . /BIN_NAME
WORKDIR /mail-core

RUN cd cmd/mail-core && CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o mail-core

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /mail-core/cmd/mail-core .
COPY --from=builder /mail-core/configs ./configs

ENTRYPOINT [ "./mail-core" ]
