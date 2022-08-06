# APIs

API security risks that must be considered, don't overexpose data, need to implement rate limits.

The back-end API does not provide all the user-friendly tips, only for the client to return the appropriate information, for example, will not return the minimum and maximum length required for the user name, but will be detailed in the API documentation, if the client-side implementers stand on the user-friendly point of view, the implementation of the client needs to be front-end verification.

Regarding the API documentation, there is no better solution for the time being, and we don't want to consider swagger at the moment.

---

## API TABLE EXAMPLE.

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
|  |  | √ | × |

## Hvx APIs V1

### Hvx Public APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| GetInstanceDetails | Get current instance details | √ | GET /public/instance |
| GetInstanceList | Get a list of instances | √ | GET /public/instances |
| CreateAccounts | Create an account | √ | POST /public/account/create |
| Authenticate | Login Account | √ | POST /auth |

Users created with just a username and password are not allowed to use the service properly, and the client developer needs to continue to prompt the user to improve the user information, such as adding the user's email address for authentication or recovering the password. When the user information is perfected, the user can use the service normally.

### Hvx Public ActivityPub APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| GetWebfinger | Get WebFinger [rfc7033] | √ | GET /.well-known/webfinger |
| GetActor | Get Actor details | √ | GET /u/{actor} |
| Inbox | Inbox | √ | POST /u/{actor}/inbox |

### Hvx Account APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| IsExist | Determine if an account exists by username | √ | × |
| Create | Create an account | √ | × |
| GetByUsername | Get account details by username | √ | GET /api/v1/account/{username} |
| Delete | Delete an account by account ID | √ | DELETE /api/v1/account |
| EditUsername | Change username | √ | PATCH /api/v1/account/username |
| EditPassword | Change password | √ | PATCH /api/v1/account/password |
| EditEmail | Editorial Email Address | √ | PATCH /api/v1/account/mail |
| Verify | Authentication with username and password |  |  |

### Hvx Actor APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| Create | Creating Actor Data | √ | × |
| GetActorsByPreferredUsername | Search for users by PreferredUsername | √ | GET /api/v1/actor/search/{preferred_username} |
| GetActorByAddress | Get Actor data by Actor address | √ | × |
| Edit | Edit Actor profile | √ | PUT /api/v1/actor |
| Delete | Delete Actor by ID | √ | × |
| GetActorByUsername | Get Actor by username | √ | × |

### Hvx Authorization APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| Verify | Verify account and return TOKEN | √ | × |
| DHRequest | Get DH key by specifying device id | √ | POST /api/v1/dh/request |
| SendDH | Send DH key | √ | POST /api/v1/dh |
| GetPrivate | Get PrivateKey | √ | GET /api/v1/dh/private |

### Hvx Device APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | -- | --- |
| IsExist | Determine if a device exists by ID | √ | × |
| Create | Create Device | √ | × |
| Get | Get device details by ID | √ | × |
| Delete | Deleting a device by ID | √ | DELETE /api/v1/device/{device_id} |
| GetDevices | Get a list of devices by account ID | √ | GET /api/v1/deice/{LIMIT} |
| DeleteDevices | Delete all devices by account ID | √ | × |

### Hvx Saved APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| Create | Create Saved | √ | POST /api/v1/saved |
| GetSaved | Get Saved by ID Details | √ | GET /api/v1/saved |
| GetSaves | Get Saved List | √ | GET /api/v1//saved/saves |
| EditSaved | Edited Saved Notes | √ | PUT /api/v1/saved |
| Delete | Delete this Saved | √ | DELETE /api/v1/saved |
| DeleteSaves | Delete all Saved | √ | DELETE /api/v1/saved/saves |

### Hvx Message APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| AccessRegister | Create a Message account | √ | POST /api/v1/message/register |
| AccessLogin | Create a Message account | √ | POST /api/v1/message/login |
| AccessDelete | Delete Message account | √ | DELETE /api/v1/message/access |

### Hvx Channel APIs

**Administrative**

| Methods | Comment | gRPC | RESTful | Authorization |
| --- | --- | --- | --- | --- |
| IsAdministrator | Determine if the channel is Channel GetAdministrator | √ | × |  |
| IsOwner | Is the channel owner | √ | × |  |
| AddAdministrator | Add Administrator | √ | POST /api/v1/channel/admin | OWNER / ADMIN |
| RemoveAdministrator | Remove Channel Administrator | √ | DELETE /api/v1/channel/admin | OWNER |
| GetAdministrators | Get all channel administrators | √ | GET /api/v1/channel/admin/{channel_id}/admins | OWNER / ADMIN |
| ExitAdministrator | Exit Administrator | √ | POST /api/v1/channel/admin/exit | ADMIN |

**Broadcast**

| Methods | Comment | gRPC | RESTful | Authorization |
| --- | --- | --- | --- | --- |
| CreateBroadcast | Create a broadcast | √ | POST /api/v1/channel/broadcast | OWNER / ADMIN |
| GetBroadcasts | Get all broadcasts in the channel | √ | GET /api/v1/channel/broadcast/{channel_id} | OWNER / ADMIN / SUBSCRIBER |
| DeleteBroadcast | Delete a broadcast | √ | DELETE /api/v1/channel/broadcast | OWNER / ADMIN |

**Channel**

| Methods | Comment | gRPC | RESTful | Authorization |
| --- | --- | --- | --- | --- |
| CreateChannel | Create Channel | √ | POST /api/v1/channel | OWNER |
| GetChannels | Get all your own and managed Channels | √ | GET /api/v1/channel/channels | OWNER /  ADMIN |
| DeleteChannel | Delete Channel | √ | DELETE /api/v1/channel | OWNER |
| EditChannel | Using the EditActor API | x | x | OWNER |
| DeleteChannels | Delete all channels | √ | x | OWNER  |

**Subscriber**

| Methods | Comment | gRPC | RESTful | Authorization |
| --- | --- | --- | --- | --- |
| IsSubscriber | Is a subscriber | x | x |  |
| AddSubscriber | Add subscribed members | √ | POST /api/v1/channel/subscriber | OWNER / ADMIN |
| RemoveSubscriber | Remove subscribed members | √ | DELETE /api/v1/channel/subscriber | OWNER / ADMIN |
| GetSubscribers | Get a list of subscribed members | √ | GET /api/v1/channel/{channel_id}/subscribers | OWNER / ADMIN |
| Subscription | Subscribe to the channel | √ | POST /api//v1/channel/subscription |  |
| Unsubscribe | Unsubscribe | √ | DELETE /api/v1/channel/unsubscribe |  |

### Hvx Article APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| Create | Create a post or status | √ | POST /api/v1/article |
| Get | Get article or status details | √ | GET /api/v1/article |
| GetArticles | Get all posts by Actor | √ | GET /api/v1/article/{LIMIT} |
| Update | Update article | √ | PUT /api/v1/article |
| Delete | Delete article | √ | DELETE /api/v1/article |
| DeleteArticles | Delete all articles | √ | DELETE /api/v1/articles |

### Hvx Notify APIs

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| Subscribe | Subscribe to notifications | √ | POST /api/v1/notify/subscribe |

### Hvx Activity APIs

**Inbox**

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| Inbox | Inbox | √ | x |
| GetInbox | Get email details by campaign ID | √ | GET /api/v1/activity/inbox |
| GetInboxes | Get all inboxes | √ | GET /api/v1/activity/inbox/{LIMIT} |
| DeleteInbox | Delete an email | √ | DELETE /api/v1/activity/inbox |
| CreateOutbox | Create an outbox message | √ | x |
| GetOutbox | Get Outbox by Activity ID | √ | GET /api/v1/activity/outbox |
| GetOutboxes | Get the outbox list | √ | GET /api/v1/activity/outbox/{LIMIT} |

**Follow**

| Methods | Comment | gRPC | RESTful |
| --- | --- | --- | --- |
| GetFollowers | Get a list of followers | √ | GET /api/v1/activity/followers |
| GetFollowings | Get on the Watch List | √ | GET /api/v1/activity/following |
| Follow | Start a concern | √ | POST /api/v1/activity/follow |
| UnFollow | Unfollow | √ | POST /api/v1/activity/unfollow |
| GetFriends | Get a list of friends | √ | GET /api/v1/activity/friends |