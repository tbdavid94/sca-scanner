#!/usr/bin/env bash

IMAGE=$1

docker run --rm -e "WORKSPACE=${PWD}" -v $PWD:/app -v $PWD/logs:/tmp "$IMAGE"
#run base no entrypoint
#docker run --rm -e "WORKSPACE=${PWD}" -v $PWD:/app -v $PWD/logs:/tmp sca-scan:base