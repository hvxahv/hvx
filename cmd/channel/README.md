# Channel
## Services that support Content publishing channel based on ActivityPub service

The service is specifically implemented in the `./internal `

`*_handler.go` For processing gRPC requests

`administrative.go` Interface to channel administrators

`broadcast.go` Published broadcast interface

`channel.go` Implementation of channel
 
`subscriber.go` Interface to Subscriber Services

Start gRPC with HTTP server.

```bash
channel run
```
The profile must be satisfied:

https://github.com/hvxahv/hvx/blob/main/conf/

