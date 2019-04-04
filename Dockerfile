# Stage build
FROM golang:1.12.1-alpine3.9 AS builder

COPY . /go/src/github.com/vaksi/safeZone
WORKDIR /go/src/github.com/vaksi/safeZone

# Download the project dependencies
RUN apk add --no-cache git mercurial && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Stage Runtime Applications
FROM alpine:latest

# Download Depedencies
RUN apk update && apk add ca-certificates bash jq curl && rm -rf /var/cache/apk/*

# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

#RUN adduser -D admin admin

ENV BUILDDIR /go/src/github.com/vaksi/safeZone

# Setting folder workdir
WORKDIR /opt/safeZone

# Copy Data App
COPY --from=builder $BUILDDIR/safeZone safeZone

ENTRYPOINT ["./safeZone"]