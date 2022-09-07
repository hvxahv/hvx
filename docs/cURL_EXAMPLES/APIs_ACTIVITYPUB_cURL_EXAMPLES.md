# ACTIVITYPUB

## GET WEBFINGER
hvxahv.disism.internal:8080/.well-known/webfinger?resource=acct:hvturingga@halfmemories.com

REQ
```shell
curl --location --request GET 'hvxahv.disism.internal:8080/.well-known/webfinger?resource=acct:hvturingga@halfmemories.com'
```
RES
```json
{
    "subject": "acct:hvturingga@halfmemories.com",
    "aliases": [
        "https://halfmemories.com/u/hvturingga"
    ],
    "links": [
        {
            "rel": "self",
            "type": "application/activity+json",
            "href": "https://halfmemories.com/u/hvturingga"
        }
    ]
}
```

## GET ACTOR
hvxahv.disism.internal:8080/u/hvturingga

REQ
```shell
curl --location --request GET 'hvxahv.disism.internal:8080/u/hvturingga'
```
RES
```json
{
    "@context": [
        "https://www.w3.org/ns/activitystreams",
        "https://w3id.org/security/v1"
    ],
    "id": "https://halfmemories.com/u/hvturingga",
    "type": "Person",
    "following": "",
    "followers": "",
    "inbox": "https://halfmemories.com/u/hvturingga/inbox",
    "outbox": "https://halfmemories.com/u/hvturingga/outbox",
    "preferredUsername": "hvturingga",
    "name": "",
    "summary": "",
    "url": "https://halfmemories.com/u/hvturingga",
    "publicKey": {
        "id": "https://halfmemories.com/u/hvturingga#main-key",
        "owner": "https://halfmemories.com/u/hvturingga",
        "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw9mcKDv+SHmSgdWMEic6\nRfRwHNOj16chn9vgka+Y32TFACl5Xoutlih6Je+LYoLmOFMlg+51lo4vnO+imCsk\nIqc+U+4Ql7N6s0fn8+x5/fYRaldnv3xV6Vb75sdq07zDF27BfXmTQ+dhmgVDqBcj\nM81MDSjHEu9KkISgTvwsuf2Pu5KQ3vot9jdZK3BMt+GeV3MJpOwpKn7OpPerXp82\nMYM39c825uc9ZnxqkhRgxL1Kw\n+JmKizaeHk8EdsmwGI09pTyFuwzexjx8QJAFsNO\nwYY0qDQrbRPOw0YQFBRCZKsci7vXxcwuFuMK+2G3SIxSEXjUq0bUq93hWjv4H2sC\nMQIDAQAB\n-----END PUBLIC KEY-----"
    }
}
```
## SEARCH ACTORS
hvxahv.disism.internal:8080/api/v1/search/hvturingga

REQ
```shell
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/search/hvturingga' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "actors": [
        {
            "id": "785503307693490177",
            "preferredUsername": "hvturingga",
            "domain": "halfmemories.com",
            "avatar": "",
            "name": "",
            "summary": "",
            "inbox": "https://halfmemories.com/u/hvturingga/inbox",
            "address": "https://halfmemories.com/u/hvturingga",
            "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw9mcKDv+SHmSgdWMEic6\nRfRwHNOj16chn9vgka+Y32TFACl5Xoutlih6Je+LYoLmOFMlg+51lo4vnO+imCsk\nIqc+U+4Ql7N6s0fn8+x5/fYRaldnv3xV6Vb75sdq07zDF27BfXmTQ+dhmgVDqBcj\nM81MDSjHEu9KkISgTvwsuf2Pu5KQ3vot9jdZK3BMt+GeV3MJpOwpKn7OpPerXp82\nMYM39c825uc9ZnxqkhRgxL1Kw\n+JmKizaeHk8EdsmwGI09pTyFuwzexjx8QJAFsNO\nwYY0qDQrbRPOw0YQFBRCZKsci7vXxcwuFuMK+2G3SIxSEXjUq0bUq93hWjv4H2sC\nMQIDAQAB\n-----END PUBLIC KEY-----",
            "actorType": "Person",
            "isRemote": "false"
        }
    ]
}
```

## EDIT ACTOR
hvxahv.disism.internal:8080/api/v1/actor/

REQ
```bash
curl --location --request PUT 'hvxahv.disism.internal:8080/api/v1/actor/' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "HVTURINGGA",
    "avatar": "https://avatars.githubusercontent.com/u/35920389?v=4",
    "summary": "bio"
}'
```
RES
```json
{
    "code": "200",
    "reply": "ok"
}
```