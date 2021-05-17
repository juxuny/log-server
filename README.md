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
    image: juxuny/log-server:v0.0.1
    entrypoint:
      - /app/logd
      - "-d=/app/log"
    restart: always
    volumes:
      - ./tmp:/app/log
    ports:
      - "40000:40000"
    
```