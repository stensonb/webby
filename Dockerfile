FROM scratch
MAINTAINER Bryan Stenson <bryan@revinate.com>
COPY webby webby
ENTRYPOINT ["/webby"]
