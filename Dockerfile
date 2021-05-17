FROM golang:1.16 as build

ENV GO111MODULE=on

COPY ./ /go/src/github.com/andreip-og/gitlab-exporter/
WORKDIR /go/src/github.com/andreip-og/gitlab-exporter/

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /bin/main

FROM alpine:3.11.3

RUN apk --no-cache add ca-certificates \
     && addgroup exporter \
     && adduser -S -G exporter exporter
ADD VERSION .
USER exporter
COPY --from=build /bin/main /bin/main
ENV LISTEN_PORT=8081
ENTRYPOINT [ "/bin/main" ]
