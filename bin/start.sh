#!/bin/bash

docker run --name fastapi-admin -d -p 8080:8080 \
  --network host \
  --restart=always \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/fastapi-admin/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  -v /data/fastapi-admin/resource/public:/app/resource/public \
  iimeta/fastapi-admin:latest
