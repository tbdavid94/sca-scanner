#!/usr/bin/env bash

docker run --rm -e "WORKSPACE=${PWD}" -v $PWD:/app -v $PWD/logs:/tmp sca-scan:dev