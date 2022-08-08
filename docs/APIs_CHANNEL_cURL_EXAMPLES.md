# CHANNEL

## CREATE CHANNELS
hvxahv.disism.internal:8080/api/v1/channel

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/channel' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "preferred_username": "hvx"
}'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## GET CHANNELS
hvxahv.disism.internal:8080/api/v1/channel/channels

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/channel/channels' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
  "code": "200",
  "channels": [
    {
      "channel_id": "785746792214626305",
      "channel": {
        "id": "785746790721323009",
        "preferredUsername": "hvx",
        "domain": "halfmemories.com",
        "avatar": "",
        "name": "",
        "summary": "",
        "inbox": "https://halfmemories.com/u/hvx/inbox",
        "address": "https://halfmemories.com/u/hvx",
        "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu5nqv6DxtvuqB48Wi/hp\nl62KSY4Ma4h0ARQ/XhZzYnqCuLl372fM+/RLWHcD21Ji+QDwjo6lg6tgf/oN1H8f\na+z0HVG848B/eZl5wxcmyGxbdS77ju48pFW/ERFsOJGAhx3oM3++2rxgLYYw9l2y\n4EQwAk/OXSOPmFDe0/4w8ghLhuKzpA8cvSmV5K5ebnajWJTpH0sO7hYoDlK/Morp\nHKKxTc963dzDPBY3pPYN8h+g0h9MzIortK3FB02Qqd64mmWguPWvD848K+xvD9Ah\n2hOosoqAA1e/CA9LK2AZf9P+lrRPnqb46tM18CyVIoRI8ez1CbeZ+gnRfscIOGiU\niwIDAQAB\n-----END PUBLIC KEY-----\n",
        "actorType": "service",
        "isRemote": "false"
      }
    }
  ]
}
```

## DELETE CHANNEL
hvxahv.disism.internal:8080/api/v1/channel

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/channel' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "channel_id": "785743169680474113"
}'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## GET ADMINISTRATORS
hvxahv.disism.internal:8080/api/v1/channel/admin/<CHANNEL_ID>

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/channel/admin/<CHANNEL_ID>' \
--header 'Authorization: Bearer <TOKEN>'
```

RES
```json
{
  "code": "200",
  "admins": [
    {
      "isOwner": true,
      "admin": {
        "id": "785518573776797697",
        "preferredUsername": "hvturingga",
        "domain": "halfmemories.com",
        "avatar": "https://avatars.githubusercontent.com/u/35920389?v=4",
        "name": "HVTURINGGA",
        "summary": "HVXAHV",
        "inbox": "https://halfmemories.com/u/hvturingga/inbox",
        "address": "https://halfmemories.com/u/hvturingga",
        "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw9mcKDv+SHmSgdWMEic6\nRfRwHNOj16chn9vgka+Y32TFACl5Xoutlih6Je+LYoLmOFMlg+51lo4vnO+imCsk\nIqc+U+4Ql7N6s0fn8+x5/fYRaldnv3xV6Vb75sdq07zDF27BfXmTQ+dhmgVDqBcj\nM81MDSjHEu9KkISgTvwsuf2Pu5KQ3vot9jdZK3BMt+GeV3MJpOwpKn7OpPerXp82\nMYM39c825uc9ZnxqkhRgxL1Kw\n+JmKizaeHk8EdsmwGI09pTyFuwzexjx8QJAFsNO\nwYY0qDQrbRPOw0YQFBRCZKsci7vXxcwuFuMK+2G3SIxSEXjUq0bUq93hWjv4H2sC\nMQIDAQAB\n-----END PUBLIC KEY-----",
        "actorType": "Person",
        "isRemote": "false"
      }
    }
  ]
}
```

## REMOVE ADMINISTRATOR
hvxahv.disism.internal:8080/api/v1/channel/admin

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/channel/admin' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "channel_id": "786088233592553473",
    "removed_id": "785747557033967617"
}'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```