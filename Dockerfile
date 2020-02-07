FROM golang AS builder
ADD . /workspace/go/src/github.com/tomgeorge/go-http-server
WORKDIR /workspace/go/src/github.com/tomgeorge/go-http-server
ENV GOPATH=/workspace/go
RUN make

FROM ubuntu
COPY --from=builder /workspace/go/src/github.com/tomgeorge/go-http-server/bin/httpserver /bin/httpserver
CMD ["/bin/httpserver"]
