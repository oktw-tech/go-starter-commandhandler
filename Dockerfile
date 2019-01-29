FROM scratch
 COPY ./go-starter-commandhandler /commandhandler
 ENTRYPOINT ["/commandhandler"]