FROM golang:1.18-alpine as builder

WORKDIR /app
COPY . .

RUN go build -o hostpath-provisioner .

FROM alpine:latest
COPY --from=builder /app/hostpath-provisioner /usr/local/bin/hostpath-provisioner

ENTRYPOINT ["/usr/local/bin/hostpath-provisioner"]
