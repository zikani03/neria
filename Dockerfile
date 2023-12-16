FROM golang:1.21-alpine as go-builder
RUN apk add --no-cache git
WORKDIR /go/neria
COPY neria-fly .
RUN go generate -x -v
RUN go build -o /bin/neria 
RUN chmod +x /bin/neria

MAINTAINER Zikani Nyirenda Mwase <zikani.nmwase@ymai.com>
FROM golang:1.21-alpine
VOLUME /opt/
VOLUME /data/
EXPOSE 8080
COPY --from=go-builder /bin/neria /neria
CMD ["/neria"]

