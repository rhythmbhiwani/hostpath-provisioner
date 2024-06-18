FROM golang:1.22-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Now copy the rest of the files
COPY . .

# Build the Go app
RUN go build -o hostpath-provisioner .

FROM alpine:latest
COPY --from=builder /app/hostpath-provisioner /usr/local/bin/hostpath-provisioner

ENTRYPOINT ["/usr/local/bin/hostpath-provisioner"]
