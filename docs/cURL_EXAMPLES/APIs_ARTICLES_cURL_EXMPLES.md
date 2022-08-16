# ARTICLE

## CREATE ARTICLE
hvxahv.disism.internal:8080/api/v1/article

REQUEST
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/article' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":  "Rebirth.",
    "summary": "DISISM ❤️ YOU",
    "article": "<p>In January 2022, disism.com took on a new lease of life and did something more meaningful to protect digital privacy and security, we launched the HVXAHV program.</p>",
    "tags":  ["disism"],
    "attachments_type": "",
    "attachments": [""],
    "to": ["785747557033967617"],
    "cc": [""],
    "state": false,
    "nsfw": false,
    "visibility": "1"
}'
```
RESPONSE
```json
{
    "code": "200",
    "status": "ok"
}
```

## CREATE STATUS
hvxahv.disism.internal:8080/api/v1/article

REQUEST
```bash
{
    "article": "<p>Sayonara Детка</p>",
    "tags":  ["Arts"],
    "attachments_type": "Video",
    "attachments": ["https://www.youtube.com/watch?v=ZNmX9NK1UVs"],
    "to": ["785747557033967617"],
    "cc": [""],
    "state": false,
    "nsfw": false,
    "visibility": "1"
}
```
RESPONSE
```json
{
    "code": "200",
    "status": "ok"
}
```

## GET ARTICLE
hvxahv.disism.internal:8080/api/v1/article/<ARTICLE_ID>

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/article/787516945347018753' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "actor": {
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
    },
    "article": {
        "id": "787516945347018753",
        "title": "Rebirth.",
        "summary": "DISISM ❤️ YOU",
        "article": "<p>In January 2022, disism.com took on a new lease of life and did something more meaningful to protect digital privacy and security, we launched the HVXAHV program.</p>",
        "tags": [
            "disism"
        ],
        "attachmentType": "",
        "attachments": [
            ""
        ],
        "to": [
            "785747557033967617"
        ],
        "cc": [
            "0"
        ],
        "state": false,
        "nsfw": false,
        "visibility": "1"
    }
}
```

## GET ARTICLES
hvxahv.disism.internal:8080/api/v1/article/articles

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/article/articles' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "data": [
        {
            "id": "787516945347018753",
            "title": "Rebirth.",
            "summary": "DISISM ❤️ YOU",
            "article": "<p>In January 2022, disism.com took on a new lease of life and did something more meaningful to protect digital privacy and security, we launched the HVXAHV program.</p>",
            "tags": [
                "disism"
            ],
            "attachmentType": "",
            "attachments": [
                ""
            ],
            "to": [
                "785747557033967617"
            ],
            "cc": [
                "0"
            ],
            "state": false,
            "nsfw": false,
            "visibility": "1"
        },
        {
            "id": "787516952497618945",
            "title": "",
            "summary": "",
            "article": "<p>Sayonara Детка</p>",
            "tags": [
                "Arts"
            ],
            "attachmentType": "Video",
            "attachments": [
                "https://www.youtube.com/watch?v=ZNmX9NK1UVs"
            ],
            "to": [
                "785747557033967617"
            ],
            "cc": [
                "0"
            ],
            "state": false,
            "nsfw": false,
            "visibility": "1"
        }
    ]
}
```

## EDIT ARTICLE
hvxahv.disism.internal:8080/api/v1/article

REQ
```bash
curl --location --request PUT 'hvxahv.disism.internal:8080/api/v1/article' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "787516945347018753",
    "title": "",
    "summary": "JUST 4 ▷ FUN, DISISM ❤️ YOU",
    "article": "",
    "tags": [""],
    "attachments_type": "",
    "attachments": [""],
    "nsfw": "",
    "visibility": ""
}'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## DELETE ARTICLE
hvxahv.disism.internal:8080/api/v1/article

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/article' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "<ARTICLE_ID>"
}'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## DELETE ALL ARTICLES
hvxahv.disism.internal:8080/api/v1/article/articles

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/article/articles' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```