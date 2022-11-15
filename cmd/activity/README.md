# Activity
## Services that support Activitypub protocol

The service is specifically implemented in the `./internal `

`*_handler.go` For processing gRPC requests

`activity_handler.go` Handle posting activities.

`friendship_handler.go` Handling buddy relationships.

`inbox_handler.go` Handle inboxes.

`outbox_handler.go` Handles outgoing mailboxes.

`/activity/*.go` The specific component codes of the various activitypub activities.

`/delivery/delivery.go` The specific component codes of the various activitypub activities.


Start gRPC with HTTP server.

```bash
activity run
```
The profile must be satisfied:

https://github.com/hvxahv/hvx/blob/main/conf/

