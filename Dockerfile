# Grab golang official alpine image and use as the base image
FROM golang:1.18.3-alpine3.16 as builder

LABEL maintainer="Derek Cleveland <clevelanddtc@gmail.com>"

WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

# Grab the latest versions of packages used by the app
ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

# Change directory of particular main.go
WORKDIR cmd/band-app-server

# Build the binary
# CGO_ENABLED=0: Disables the use of CGO. While it should be disabled while cross compiling its not always the case
# GOOS=linux: target operating system linux
# ldfflags: -w turns off DWARF debugging information, -s turns off generation of the Go symbol table
# -a: force rebuilding of packages that are already up-to-date. Slows build time but know everything is built properly
# -o: forces build to write the executable to the named output or directory
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -o /go/bin/band-app-server .

# Build 2 make a small image
FROM alpine:3.16

RUN apk update && \
    apk add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

# Copy the binary from builder
COPY --from=builder /go/bin/band-app-server /go/bin/band-app-server

# Document that the container uses port 8080 - this is just for documentation
EXPOSE 8080

# Run the binary
CMD ["/go/bin/band-app-server"]