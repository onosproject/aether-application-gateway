FROM onosproject/golang-build:v0.6.10 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /go/src/github.com/onosproject/aether-application-gateway/
COPY cmd/ /go/src/github.com/onosproject/aether-application-gateway/cmd/
COPY cmd/ /go/src/github.com/onosproject/aether-application-gateway/internal/
COPY pkg/ /go/src/github.com/onosproject/aether-application-gateway/pkg/
COPY vendor/ /go/src/github.com/onosproject/aether-application-gateway/vendor/

RUN cd /go/src/github.com/onosproject/aether-application-gateway && go build -mod=vendor ./cmd/aether-application-gateway

FROM alpine:3.12
RUN apk add libc6-compat

USER nobody

COPY --from=build /go/src/github.com/onosproject/aether-application-gateway/build/_output/aether-application-gateway /usr/local/bin/aether-application-gateway
COPY api/ /

ENTRYPOINT ["aether-application-gateway"]
