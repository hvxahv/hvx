# MESSAGE

## MATRIX ACCESS REGISTER
hvxahv.disism.internal:8080/api/v1/message/access/register

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/message/access/register' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "<USERNAME>",
    "password": "<PASSWORD>"
}'
```

RES
```json
{
  "code": "200",
  "access_token": "<TOKEN>",
  "user_id": "@<USERNAME>:<MATRIX_SERVICES>"
}
```