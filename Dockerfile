FROM golang:1.16-alpine AS builder

ARG srcpath="/build/hostpath-provisioner"

RUN apk --no-cache add git && \
    mkdir -p "$srcpath"

ADD . "$srcpath"

RUN cd "$srcpath" && \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    go build -a -ldflags '-extldflags "-static"' -o /hostpath-provisioner

FROM scratch

COPY --from=builder /hostpath-provisioner /hostpath-provisioner

CMD ["/hostpath-provisioner"]
