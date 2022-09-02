## FS 

## AVATAR
hvxahv.disism.internal:8080/api/v1/fs/avatar

Accept PNG/JPEG format images as avatars.

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/fs/avatar' \
--header 'Authorization: Bearer <TOKEN>' \
--form 'avatar=@"<FILE_DIR>/<IMAGE_NAME>"'
```
RES
```json
{
  "avatar": "http://hvxahv.disism.internal:9000/avatar/<FILE_NAME>.<IMAGE>",
  "code": "200",
  "name": "<FILE_NAME>"
}
```

## ATTACH
hvxahv.disism.internal:8080/api/v1/fs/attach

REQ
```bash
curl --location --request POST 'hvxahv.disism.internal:8080/api/v1/fs/attach' \
--header 'Authorization: Bearer <TOKEN>' \
--form 'attach=@"<FILE_DIR>/<FILE_NAME>"' \
--form 'attach=@"<FILE_DIR>/<FILE_NAME>"'
```
RES
```json
{
    "attas": [
        {
            "name": "<FILE_NAME>",
            "address": "<FILE_ADDRESS>"
        },
        {
            "name": "<FILE_NAME>",
            "address": "<FILE_ADDRESS>"
        },
    ],
    "code": "200"
}
```
## DELETE
hvxahv.disism.internal:8080/api/v1/fs/source

REQ
```bash
curl --location --request DELETE 'hvxahv.disism.internal:8080/api/v1/fs/source' \
--header 'Authorization: Bearer <TOKEN>' \
--form 'name="<FILE_NAME>"'
```
RES
```json
{
    "code": "200",
    "status": "ok"
}
```

## GET FILE ADDRESS
hvxahv.disism.internal:8080/api/v1/fs/address/<FILE_NAME>

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/fs/address/<FILE_NAME>' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
    "address": "<ADDRESS>",
    "code": "200",
    "name": "<FILE_NAME>"
}
```