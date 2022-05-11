FROM ubuntu:latest
ADD dgate /
ADD dgate.yaml /
CMD ["/dgate"]

