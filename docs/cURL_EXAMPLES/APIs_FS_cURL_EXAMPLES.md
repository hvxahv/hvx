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
    "avatar": "http://hvxahv.disism.internal:9000/avatar/cd4c937c-661c-4b6e-80b4-58afaf1b7fcc.jpeg",
    "code": "200"
}
```