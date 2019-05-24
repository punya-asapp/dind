#!/bin/bash

cp $DOCKER_CERT_PATH docker.crt
export DOCKER_CERT_PATH=$(pwd)/docker.crt
export DOCKER_TLS_VERIFY=1

go test
