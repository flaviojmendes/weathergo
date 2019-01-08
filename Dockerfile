############################
# STEP 1 build executable binary
############################

FROM golang:1.11-alpine as builder
COPY . /go/src/flaviojmendes/weathergo
ENV GO111MODULE=on
WORKDIR /go/src/flaviojmendes/weathergo
RUN apk -U add git build-base && \
    mkdir -p /build && \
    go build  -ldflags '-extldflags "-static"' -o /build/weathergo


############################
# STEP 2 build a small image
############################
FROM alpine:3.7
RUN apk -U add ca-certificates curl && rm -rf /var/cache/apk/*

WORKDIR /opt
COPY --from=builder /build/weathergo .
ENV GIN_MODE=release
CMD ["./weathergo"]
