# DEVICE

## GET DEVICES 
hvxahv.disism.internal:8080/api/v1/device/devices

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/device/devices' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "code": "200",
    "devices": [
        {
            "id": "785538302981144577",
            "accountId": "785518574097694721",
            "device": "chrome",
            "createdAt": "2022-08-06 22:54:25"
        },
        {
            "id": "785541187591340033",
            "accountId": "785518574097694721",
            "device": "chrome",
            "createdAt": "2022-08-06 23:09:05"
        }
    ]
}
```

## DELETE DEVICE

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/device' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "device_id": "785553656326389761"
}'
```
RES
```json
{
    "code": "200",
    "reply": "ok"
}
```