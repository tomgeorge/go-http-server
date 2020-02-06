FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-http-server"]
COPY ./bin/ /