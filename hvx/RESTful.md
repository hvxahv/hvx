# Public RESTful API DOCS.

APIs that do not require authentication:

## health

API health check, sending ping will return pong

```
curl --location --request GET'localhost:8088/ping'
```

Response:

```bash
{
     "message": "pong"
}
```

## Account Service

### New account

Receiving user's username and password

```bash
curl --location --request POST'localhost:8088/account/new' \
--form'username="hvturinggas"' \
--form'password="hvxahv123"'
```

Response:

```bash
{
     "code": 201,
     "message": "NEW ACCOUNT OK!"
}
```

Returning 201 means that the user was created successfully, 202 means that the user already exists, and 500 means that an error occurred when the server created the user.

However, users created only with a user name and password cannot be used normally. The developer of the client needs to continue to prompt the user to complete the user information, such as adding the user's mailbox for verification or password recovery.