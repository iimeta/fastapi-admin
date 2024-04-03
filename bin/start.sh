#!/bin/bash

docker run -d \
  --network host \
  --restart=always \
  -p 8080:8080 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/fastapi-admin/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  -v /data/fastapi-admin/resource/public:/app/resource/public \
  --name fastapi-admin \
  iimeta/fastapi-admin:0.1.1
