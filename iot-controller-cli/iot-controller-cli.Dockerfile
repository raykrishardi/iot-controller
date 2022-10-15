ARG ALPINE_VERSION=3.16.2
FROM alpine:${ALPINE_VERSION}
LABEL maintainer="Ray Krishardi Layadi <raykrishardi@gmail.com>"

COPY iot /

ENTRYPOINT ["/iot"]