# Device
## Services that support Device management for login
The service is specifically implemented in the `./internal `

`handler.go` For processing gRPC requests

`device.go` Specific implementation of device management interface

Start gRPC with HTTP server.

```bash
device run
```
The profile must be satisfied:

https://github.com/hvxahv/hvx/blob/main/conf/

