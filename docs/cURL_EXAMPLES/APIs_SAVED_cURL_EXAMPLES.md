# SAVED

## CREATE SAVED
hvxahv.disism.internal:8080/api/v1/saved

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/saved' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "YUI",
    "comment": "YUI IMAGES",
    "cid": "QmbmC4H63gEDXivLcxUL9h86ZCiLb3YDmTk1inkpMDtvqC",
    "file_type": "images/png",
    "is_private": false
}'
```
RES
```json
{
    "code": "200",
    "reply": "ok"
}
```

## GET SAVED
hvxahv.disism.internal:8080/api/v1/saved/<SAVED_ID>

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/saved/<SAVED_ID>' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "id": "787742204368257025",
    "name": "YUI",
    "comment": "YUI IMAGES",
    "cid": "QmbmC4H63gEDXivLcxUL9h86ZCiLb3YDmTk1inkpMDtvqC",
    "types": "images/png",
    "createdAt": "2022-08-14 17:44:02"
}
```

## GET SAVES
hvxahv.disism.internal:8080/api/v1/saved/saves

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/saved/saves' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "saves": [
        {
            "id": "787742204368257025",
            "name": "YUI",
            "comment": "YUI IMAGES",
            "cid": "QmbmC4H63gEDXivLcxUL9h86ZCiLb3YDmTk1inkpMDtvqC",
            "types": "images/png",
            "createdAt": "2022-08-14 17:44:02"
        }
    ]
}
```

## DELETE SAVED
hvxahv.disism.internal:8080/api/v1/saved

REQ
```bash
curl --location --request PUT 'hvxahv.disism.internal:8080/api/v1/saved' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": <SAVED_ID>,
    "name": "KOBAYASHI YUI"
    "": ""
}'
```
RES
```json
{
    "code": "200",
    "reply": "ok"
}
```

## DELETE SAVES
hvxahv.disism.internal:8080/api/v1/saved/saves

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/saved/saves' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "reply": "ok"
}
```