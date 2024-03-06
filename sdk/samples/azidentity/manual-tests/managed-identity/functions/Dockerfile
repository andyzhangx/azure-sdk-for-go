FROM mcr.microsoft.com/mirror/docker/library/alpine:3.16
RUN apk add gcompat
COPY functions .
RUN chmod +x functions
ENTRYPOINT ["./functions"]
