#!/bin/sh
curl -H 'Content-Type: application/json' --data '{"build": true}' -X POST https://cloud.docker.com/api/build/v1/source/$DOCKERHUB_SECRET/trigger/$DOCKERHUB_TOKEN/call/