FROM cgr.dev/chainguard/static:latest

WORKDIR /

COPY knox /main

ENTRYPOINT ["/main"]