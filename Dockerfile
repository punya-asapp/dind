FROM circleci/golang:1.12.5
ADD docker.crt go.mod go.sum dind_test.go .
ENV DOCKER_TLS_VERIFY=1
ENV DOCKER_CERT_PATH=/go/docker.crt
CMD go test
