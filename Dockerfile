FROM golang:latest

EXPOSE 8080

ADD src/main /usr/local/bin/
ADD src/data /data
ADD src/web /web
ENTRYPOINT exec /usr/local/bin/main
