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
            "id": "785717376682754049",
            "preferredUsername": "",
            "domain": "halfmemories.com",
            "avatar": "",
            "name": "",
            "summary": "",
            "inbox": "https://halfmemories.com/u//inbox",
            "address": "https://halfmemories.com/u/",
            "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw0ZaICGxqhA7Xv+iae28\ngnmWQ3AZjkhfWWRsDGzieFoAU//6KUniPX6+TP/XjfaIeQdX2mI8CxRY/iRwy3Sr\nNucNnH0d+Gjc1nEanGJex57r4hFYBKBcbt2brJCGyse8x4Mv4ClvnNLx5l/8pjx1\nmmR8WVnXeCh0flMvPEQUvTXFno7f7vZWPblsQm55Pn/kOY/o4X7+LE3sFp7eZDii\nLueN1BoOUZyLt1Tk845fv1wCk5gdYYNYhcLFv1m2ViGLzZa91Pn63+K7pnlPTqLi\nUTXYOIdT+Obasf1LxwngXmy1O2kdWsuq1+YuX8rnR08rf+njhtWLo19bjCcrSKdm\nbwIDAQAB\n-----END PUBLIC KEY-----\n",
            "actorType": "service",
            "isRemote": "false"
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