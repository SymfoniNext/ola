FROM alpine:3.4
ADD ola /bin/ola
ENTRYPOINT ["/bin/ola"]
