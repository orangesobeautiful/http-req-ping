ARG GO_VERSION=1.21.5
ARG BUILD_IMAGE=golang:$GO_VERSION-alpine3.18

FROM --platform=$BUILDPLATFORM $BUILD_IMAGE as builder

WORKDIR /app/

RUN go version

# start build
ARG TARGETOS TARGETARCH

COPY . /app/

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -ldflags="-s" -o server

FROM gcr.io/distroless/static

WORKDIR /app/

COPY --from=builder /app/server /app/server

CMD ["/app/server"]