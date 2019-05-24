ARG DOCKER_CERT_PATH
FROM circleci/golang:1.12.5
ADD $DOCKER_CERT_PATH /docker-certs
ADD go.mod go.sum dind_test.go /go/
ENV DOCKER_TLS_VERIFY=1
ENV DOCKER_CERT_PATH=/docker-certs
CMD go test
