Log-Server
============

## Installation

```yaml
version: "3.5"
x-default: &default
  logging:
    options:
      max-size: "5M"
      max-file: "5"
services:
  srv:
    image: juxuny/log-server:v0.0.6
    entrypoint:
      - /app/logd
      - "-d=/app/log"
    restart: always
    volumes:
      - ./tmp:/app/log
    ports:
      - "40000:40000"
    
```

## Install `glog`

```bash
GOPROXY=https://goproxy.cn GO111MODULE=on go install github.com/juxuny/log-server/cmd/glog@latest
```

## Other

```shell
alias log-api="glog --prefix=api --dir=${HOME}/log-server/log --expr "
alias log-api-admin="glog --prefix=api-admin --dir=${HOME}/log-server/log --expr "
alias log-api-cron="glog --prefix=api-cron --dir=${HOME}/log-server/log --expr "
alias tail-api="tail -f ${HOME}/log-server/log/api_$(date +"%Y%m%d_%H").log"
alias tail-api-admin="tail -f ${HOME}/log-server/log/api-admin_$(date +"%Y%m%d_%H").log"
alias tail-api-cron="tail -f ${HOME}/log-server/log/api-cron_$(date +"%Y%m%d_%H").log"
```
