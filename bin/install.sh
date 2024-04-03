#!/bin/bash

docker pull iimeta/fastapi-admin:0.1.1

mkdir -p /data/fastapi-admin/manifest/config
mkdir -p /data/fastapi-admin/resource/public

wget -P /data/fastapi-admin/manifest/config https://github.com/iimeta/fastapi-admin/raw/docker/manifest/config/config.yaml
wget https://github.com/iimeta/fastapi-admin/raw/docker/bin/start.sh
