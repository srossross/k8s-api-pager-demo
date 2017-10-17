FROM alpine
COPY ./k8s-test-controller-linux-amd64  /test-controller
CMD ["/test-controller"]
