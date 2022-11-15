# Actor
## Services that support Activitypub Actor protocol

The service is specifically implemented in the `./internal `

`handler.go` For processing gRPC requests

`actor.go` For a detailed implementation of the activitypub actor, please see the interface.

Start gRPC with HTTP server.

```bash
actor run
```
The profile must be satisfied:

https://github.com/hvxahv/hvx/blob/main/conf/

