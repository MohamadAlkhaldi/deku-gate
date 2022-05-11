FROM alpine:latest
ADD deku-gate /
ADD deku-gate.yaml /
CMD ["/deku-gate"]