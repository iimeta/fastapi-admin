#!/bin/bash

docker run --name fastapi-admin -d -p 8080:8080 \
  --network host \
  --restart=always \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/fastapi-admin:/app \
  iimeta/fastapi-admin:latest
