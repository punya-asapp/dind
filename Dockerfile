FROM circleci/golang:1.12.5
ADD test.sh docker.crt go.mod go.sum dind_test.go .
CMD ./test.sh
