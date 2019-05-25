#!/bin/bash

export DEBIAN_FRONTEND=noninteractive
sudo add-apt-repository -yq ppa:longsleep/golang-backports
sudo apt-get -yq update
sudo apt-get -yq install golang-go
