Place the following configuration files in the user's default environment directory and set the name `.hvxahv.yaml`

.hvxahv.yaml


```bigquery
---
localhost: ""

http_port: ""

author: ""

token_signed: ""

db:
  host: ""
  port: ""
  user: ""
  password: ""
  dbName: ""
  sslMode: ""

redis:
  host: ""
  port: ""
  password: ""

mongo:
  address: ""
  username: ""
  password: ""
  name: ""

bot:
  tg_dev_id: ""
  tg_token: ""

oos:
  minio:
    addr: ""
    accessKeyID: ""
    secretAccessKey: ""
    useSSL: false

microservices:
  accounts:
    host: ""
    port: ""

```