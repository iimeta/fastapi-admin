#!/bin/bash
cd `dirname $0`
cd ../

docker build -f ./bin/Dockerfile -t iimeta/fastapi-admin:1.4.0 .
