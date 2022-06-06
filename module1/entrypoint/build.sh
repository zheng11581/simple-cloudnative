#!/bin/sh
export BUILD_TYPE=shell
docker build -t goharbor.com/demo/discovery-service:$BUILD_TYPE . -f Dockerfile-${BUILD_TYPE}