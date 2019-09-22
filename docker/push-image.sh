#!/bin/bash

IMAGE=skybig/pulsar-operator:latest

# build operator
cd .. && make build && cd ./docker

# get pulsar oprator bin
cp ../bin/pulsar-operator .

echo "[START] build pulsar operator images"

# build docker image
docker build --tag "${IMAGE}" .

# push docker image
docker push "${IMAGE}"

echo "[END] build pulsar operator images"

# remove pulsar operator bin
rm -f ./pulsar-operator
