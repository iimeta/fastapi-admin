<div align=center>

<img src="https://www.fastapi.ai/logo.png" width="128"/>

# 智元 Fast API Admin

[![Release](https://img.shields.io/github/v/release/iimeta/fastapi-admin?color=blue)](https://github.com/iimeta/fastapi-admin/releases)
[![License](https://img.shields.io/static/v1?label=license&message=Apache-2.0&color=green)](https://github.com/iimeta/fastapi-admin/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/iimeta/fastapi-admin)](https://goreportcard.com/report/github.com/iimeta/fastapi-admin)

</div>

企业级 LLM API 快速集成系统，支持DeepSeek、OpenAI、Azure、文心一言、讯飞星火、通义千问、智谱GLM、豆包、Gemini、Anthropic Claude以及OpenAI格式的模型等，简洁的页面风格，轻量高效且稳定，支持Docker一键部署。业务系统只需要按照统一API标准，对接一次的开发工作量，即可无缝对接N个大模型，无需考虑N个大模型背后的各种复杂逻辑等等，可大大降低开发和维护成本...

## 手动Docker部署
0. 需提前安装好Redis和MongoDB以及Docker服务
1. 拉取镜像: `docker pull iimeta/fastapi-admin:latest`
2. 创建配置文件目录: `mkdir -p /data/fastapi-admin/manifest/config`
3. 创建公开文件目录: `mkdir -p /data/fastapi-admin/resource/public`
4. 进入配置文件目录: `cd /data/fastapi-admin/manifest/config`
5. 下载配置文件: `wget https://github.com/iimeta/fastapi-admin/raw/docker/manifest/config/config.yaml`
6. 修改配置文件: `vi config.yaml` 按 config.yaml 中配置项的注释, 修改对应参数
7. 启动服务: 
```shell
docker run --name fastapi-admin -d -p 8080:8080 \
  --network host \
  --restart=always \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/fastapi-admin:/app \
  iimeta/fastapi-admin:latest
```
8. 用户端访问: http://ip:8080/login 如: http://127.0.0.1:8080/login
9. 管理端访问: http://ip:8080/admin 如: http://127.0.0.1:8080/admin
10. 默认管理员账号密码: 无默认账号和密码, 当首次登录时输入的账号和密码将自动创建为管理员的账号和密码

## 自动Docker部署
0. 需提前安装好Redis和MongoDB以及Docker服务
1. 下载一键部署脚本: `wget -O install.sh https://github.com/iimeta/fastapi-admin/raw/docker/bin/install.sh && sh install.sh`
2. 修改配置文件: `vi /data/fastapi-admin/manifest/config/config.yaml` 按 config.yaml 中配置项的注释, 修改对应参数
3. 启动服务: `chmod +777 start.sh && ./start.sh`
4. 用户端访问: http://ip:8080/login 如: http://127.0.0.1:8080/login
5. 管理端访问: http://ip:8080/admin 如: http://127.0.0.1:8080/admin
6. 默认管理员账号密码: 无默认账号和密码, 当首次登录时输入的账号和密码将自动创建为管理员的账号和密码

## 备用国内镜像地址
```shell
docker pull registry.cn-hangzhou.aliyuncs.com/iimeta/fastapi-admin:latest
```
```shell
docker run --name fastapi-admin -d -p 8080:8080 \
  --network host \
  --restart=always \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/fastapi-admin:/app \
  registry.cn-hangzhou.aliyuncs.com/iimeta/fastapi-admin:latest
```


![Star History Chart](https://api.star-history.com/svg?repos=iimeta/fastapi-admin&type=Date)
