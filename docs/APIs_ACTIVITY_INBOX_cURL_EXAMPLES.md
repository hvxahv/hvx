# ACTIVITY INBOX

## GET INBOXES
hvxahv.disism.internal:8080/api/v1/activity/inboxes

REQ
```bash
curl --location --request GET 'hvxahv.disism.internal:8080/api/v1/activity/inboxes' \
--header 'Authorization: Bearer <TOKEN>'
```
RES
```json
{
  "code": "200",
  "inboxes": [
    {
      "id": "787795545018138625",
      "receiver_id": "785518573776797697",
      "sender_addr": "https://mstdn.social/users/hvturingga",
      "activity_id": "https://mstdn.social/883827e8-a051-4536-9086-5fbd22018c27",
      "activity_type": "Follow",
      "activity_body": "{\"@context\":\"https://www.w3.org/ns/activitystreams\",\"id\":\"https://mstdn.social/883827e8-a051-4536-9086-5fbd22018c27\",\"type\":\"Follow\",\"actor\":\"https://mstdn.social/users/hvturingga\",\"object\":\"https://halfmemories.com/u/hvturingga\"}"
    }
  ]
}
```