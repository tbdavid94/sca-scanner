#!/usr/bin/env bash
DOCKER_CMD=docker
if command -v podman >/dev/null 2>&1; then
    DOCKER_CMD=podman
fi
isort **/*.py
python3 -m black lib
python3 -m black scan
flake8 . --count --select=E9,F63,F7,F82 --show-source --statistics
flake8 . --count --exit-zero --statistics

$DOCKER_CMD build -t tbdavid94/sca-scanner -f Dockerfile .
$DOCKER_CMD build -t tbdavid94/sca-scanner:base -f ci/base.Dockerfile .
$DOCKER_CMD build -t tbdavid94/sca-scanner:java -f ci/java.Dockerfile .
$DOCKER_CMD build -t tbdavid94/sca-scanner:slim -f ci/dynamic-lang.Dockerfile .
$DOCKER_CMD build -t tbdavid94/sca-scanner:csharp -f ci/csharp.Dockerfile .
$DOCKER_CMD build -t tbdavid94/sca-scanner:oss -f ci/oss.Dockerfile .
