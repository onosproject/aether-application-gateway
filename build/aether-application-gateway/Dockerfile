FROM onosproject/golang-build:v1.0.0 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /go/src/github.com/onosproject/aether-application-gateway/
COPY cmd/ /go/src/github.com/onosproject/aether-application-gateway/cmd/
COPY internal/ /go/src/github.com/onosproject/aether-application-gateway/internal/
COPY pkg/ /go/src/github.com/onosproject/aether-application-gateway/pkg/
COPY vendor/ /go/src/github.com/onosproject/aether-application-gateway/vendor/

RUN cd /go/src/github.com/onosproject/aether-application-gateway && GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}

FROM alpine:3.12
RUN apk add libc6-compat

USER nobody

COPY --from=build /go/src/github.com/onosproject/aether-application-gateway/build/_output/aether-application-gateway /usr/local/bin/aether-application-gateway
COPY api/ /

ENTRYPOINT ["aether-application-gateway"]
