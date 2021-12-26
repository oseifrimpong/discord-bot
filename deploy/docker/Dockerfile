FROM golang:1.17 AS builder
LABEL key="keyspecs"
WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:3
# WORKDIR /bot
COPY --from=builder /build/discussion-bot .
COPY ./deploy/scripts/app/init.sh /

RUN chmod +x /init.sh
ENTRYPOINT ["/init.sh"]


CMD /discussion-bot