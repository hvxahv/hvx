# ACCOUNT

## EDIT USERNAME
hvxahv.disism.internal:8080/api/v1/account/username

REQ
```bash
curl --location --request PATCH 'hvxahv.disism.internal:8080/api/v1/account/username' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "hvxahv"
}'
```
RES
```json
{
  "code": "200",
  "reply": "ok"
}
```

## EDIT PASSWORD
hvxahv.disism.internal:8080/api/v1/account/password
REQ

```bash
curl --location --request PATCH 'hvxahv.disism.internal:8080/api/v1/account/password' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "hvturingga",
    "password": "hvx123",
    "new_password": "hvxahv123"
}'
```
RES
```json
{
  "code": "200",
  "reply": "ok"
}
```

## EDIT EMAIL
hvxahv.disism.internal:8080/api/v1/account/mail

REQ

```bash
curl --location --request PATCH 'hvxahv.disism.internal:8080/api/v1/account/mail' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mail": "x@disism.com"
}'
```

RES
```json
{
    "code": "200",
    "status": "ok"
}
```