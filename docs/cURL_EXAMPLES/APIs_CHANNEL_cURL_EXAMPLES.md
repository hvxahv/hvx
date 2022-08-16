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
    "channel_id": <CHANNEL_ID>,
    "removed_id": <REMOVED_ACTOR_ID>
}'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## CREATE BROADCAST
hvxahv.disism.internal:8080/api/v1/channel/broadcast

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/channel/broadcast' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "channel_id": <CHANNEL_ID>,
    "article_id": <ARTICLE_ID>
}'
```

RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## GET BROADCASTS
hvxahv.disism.internal:8080/api/v1/channel/broadcast/<CHANNEL_ID>

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/channel/broadcast/786088233592553473' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "broadcasts": [
        {
            "id": "787532659092881409",
            "channelId": "786088233592553473",
            "adminId": "785518573776797697",
            "cid": "b8a46ad4-dd7b-4670-9016-ff6c745d83c1"
        }
    ]
}
```

## DELETE BROADCAST
hvxahv.disism.internal:8080/api/v1/channel/broadcast

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/channel/broadcast' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "channel_id" : <CHANNEL_ID>",
    "broadcast_id": <BROADCAST_ID>
}'
```
RES
```json
{
    "code": "200",
    "reply": "ok"
}
```

## ADD SUBSCRIBER
hvxahv.disism.internal:8080/api/v1/channel/subscriber

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/channel/subscriber' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "channel_id": <CHANNEL_ID>,
    "subscriber_id": <SUBSCRIBER_ID>
}'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## REMOVE SUBSCRIBER
hvxahv.disism.internal:8080/api/v1/channel/subscriber

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/channel/subscriber' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "channel_id": <CHANNEL_ID>,
    "removed_id": <REMOVED_ID>
}'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## GET SUBSCRIBERS
hvxahv.disism.internal:8080/api/v1/channel/<CHANNEL_ID>/subscribers

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/channel/<CHANNEL_ID>/subscribers' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "subscriber": [
        {
            "id": "785518573776797697",
            "preferredUsername": "hvturingga",
            "domain": "halfmemories.com",
            "avatar": "https://avatars.githubusercontent.com/u/35920389?v=4",
            "name": "HVTURINGGA",
            "summary": "HVX",
            "inbox": "https://halfmemories.com/u/hvturingga/inbox",
            "address": "https://halfmemories.com/u/hvturingga",
            "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw9mcKDv+SHmSgdWMEic6\nRfRwHNOj16chn9vgka+Y32TFACl5Xoutlih6Je+LYoLmOFMlg+51lo4vnO+imCsk\nIqc+U+4Ql7N6s0fn8+x5/fYRaldnv3xV6Vb75sdq07zDF27BfXmTQ+dhmgVDqBcj\nM81MDSjHEu9KkISgTvwsuf2Pu5KQ3vot9jdZK3BMt+GeV3MJpOwpKn7OpPerXp82\nMYM39c825uc9ZnxqkhRgxL1Kw\n+JmKizaeHk8EdsmwGI09pTyFuwzexjx8QJAFsNO\nwYY0qDQrbRPOw0YQFBRCZKsci7vXxcwuFuMK+2G3SIxSEXjUq0bUq93hWjv4H2sC\nMQIDAQAB\n-----END PUBLIC KEY-----",
            "actorType": "Person",
            "isRemote": "false"
        },
        {
            "id": "785747724097224705",
            "preferredUsername": "karma",
            "domain": "halfmemories.com",
            "avatar": "",
            "name": "",
            "summary": "",
            "inbox": "https://halfmemories.com/u/karma/inbox",
            "address": "https://halfmemories.com/u/karma",
            "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw9mcKDv+SHmSgdWMEic6\nRfRwHNOj16chn9vgka+Y32TFACl5Xoutlih6Je+LYoLmOFMlg+51lo4vnO+imCsk\nIqc+U+4Ql7N6s0fn8+x5/fYRaldnv3xV6Vb75sdq07zDF27BfXmTQ+dhmgVDqBcj\nM81MDSjHEu9KkISgTvwsuf2Pu5KQ3vot9jdZK3BMt+GeV3MJpOwpKn7OpPerXp82\nMYM39c825uc9ZnxqkhRgxL1Kw\n+JmKizaeHk8EdsmwGI09pTyFuwzexjx8QJAFsNO\nwYY0qDQrbRPOw0YQFBRCZKsci7vXxcwuFuMK+2G3SIxSEXjUq0bUq93hWjv4H2sC\nMQIDAQAB\n-----END PUBLIC KEY-----",
            "actorType": "Person",
            "isRemote": "false"
        }
    ]
}
```

## SUBSCRIPTION
hvxahv.disism.internal:8080/api/v1/channel/subscription

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/channel/subscription' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "channel_id": <CHANNEL_ID>
}'
```
RES
```json
{
    "code": "200",
    "reply": "ok"
}
```

## UNSUBSCRIPTION
hvxahv.disism.internal:8080/api/v1/channel/unsubscription

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/channel/unsubscribe' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "channel_id": <CHANNEL_ID>"
}'
```
RES
```json
{
    "code": "200",
    "reply": "ok"
}
```