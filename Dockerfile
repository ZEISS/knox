# https://goreleaser.com/docker/

FROM gcr.io/distroless/static:nonroot

ARG TARGETPLATFORM

WORKDIR /
COPY $TARGETPLATFORM/knox /main

USER 65532:65532

CMD ["/main"]
