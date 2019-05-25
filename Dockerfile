ARG DOCKER_CERT_PATH
FROM circleci/golang:1.12.5
RUN sudo mkdir /proj
WORKDIR /proj
ADD go.mod go.sum dind_test.go /proj/
CMD go test
