FROM cgr.dev/chainguard/static:latest

WORKDIR /

COPY main /main

ENTRYPOINT ["/main"]