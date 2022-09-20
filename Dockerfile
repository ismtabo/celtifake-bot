FROM golang:1.18.1-alpine
RUN apk add --no-cache git make bash gcc docker-cli curl
WORKDIR /src
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download -x

COPY . /src
# Version
ARG PRODUCT_VERSION
ARG PRODUCT_REVISION
# Build the application
RUN make build-bin build-config

WORKDIR /src/build/bin
ENTRYPOINT ["./bot"]
