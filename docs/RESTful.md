# RESTful API DOCS

## SIGNUP
### REQ
```shell
curl --location --request POST 'http://localhost:8088/signup' \
--form 'username="<USERNAME>"' \
--form 'password="<PASSWORD>"' \
--form 'mail="<E-MAIL>"' \
--form 'public_key="<PUBLIC_KEY>"'
```
### RESP
ok:
```shell
{
    "code": "200",
    "message": "ok"
}
```
error:
```shell
{
    "code": "202",
    "message": "THE_USERNAME_OR_MAIL_ALREADY_EXISTS"
}
```
