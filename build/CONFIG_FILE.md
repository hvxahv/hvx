Place the configuration file in the user directory and set the name `.hvxahv.yaml`

.hvxahv.yaml


```
---
author: "disism"
name: "github.com/hvxahv/hvxahv"
version: ""

# Set domain.
localhost: "domain"

# IPFS api address.
ipfs_addr: "address"

# IPFS Gateway
ipfs_gateway: "address/ipfs/"

# Key used for signed.
token_signed: "jwt_key_hvxahv.half_memories.com"

# Set TOKEN expiration time, days.
token_expired: 60

# Matrix server
matrix:
  addr: "address"

cockroach:
  host: "address"
  port: "26257"
  user: ""
  password: ""
  dbName: ""
  sslMode: "disable"
  timeZone: "Asia/Shanghai"

redis:
  host: "address"
  port: "6379"
  password: ""

bot:
# Telegram Bot Setting
  tg_dev_id: "bot_id"
  tg_token: ""


minio:
  addr: "address"
  accessKeyID: ""
  secretAccessKey: ""
  useSSL: false
  location: "ap-northeast-3"

consul:
  address: "address"

microservices:
  # Provide RESTful API services for external access.
  hvx:
    version: "0.0.1"
    host: "hvxahv.disism.internal"
    port: "8088"

  # Account service.
  account:
    version: "0.0.1"
    host: "hvxahv.disism.internal"
    port: "7041"

  device:
    version: "0.0.1"
    host: "hvxahv.disism.internal"
    port: "7042"

  notify:
    version: "0.0.1"
    host: "hvxahv.disism.internal"
    port: "7043"

  saved:
    version: "0.0.1"
    host: "hvxahv.disism.internal"
    port: "7044"

  message:
    version: "0.0.1"
    host: "hvxahv.disism.internal"
    port: "7045"

```
