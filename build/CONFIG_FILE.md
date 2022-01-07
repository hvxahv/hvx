Place the configuration file in the user directory and set the name `.hvxahv.yaml`

.hvxahv.yaml


```
---
# Author settings.
author: "disism.com"

# Go mod repo
name: "github.com/hvxahv/hvxahv"
version: "0.0.1"

# Set domain.
localhost: "example.bar"

# IPFS api address.
ipfs_addr: ""
# IPFS Gateway
ipfs_gateway: ""

# Key used for signed.
token_signed: "jwt_key_hvxahv.foo.bar"

# Set TOKEN expiration time, days.
token_expired: 60

# Matrix server
matrix:
  addr: ""

cockroach:
  host: ""
  port: ""
  user: ""
  password: ""
  dbName: ""
  sslMode: ""
  timeZone: ""

redis:
  host: ""
  port: ""
  password: ""

bot:
  tg_dev_id: ""
  tg_token: ""


minio:
  addr: ""
  accessKeyID: ""
  secretAccessKey: ""
  useSSL: true
  location: ""

consul:
  address: ""

microservices:
  # Provide RESTful API services for external access.
  hvx:
    version: "0.0.1"
    host: "localhost"
    port: "8088"

  # Account service.
  accounts:
    version: "0.0.1"
    host: "localhost"
    port: "7041"

  # Article publishing channel service based on IPFS protocol.
  channel:
    version: "0.0.1"
    host: "localhost"
    port: "7141"

  # Instant messaging service based on matrix.org
  message:
    version: "0.0.1"
    host: ""
    port: ""
```
