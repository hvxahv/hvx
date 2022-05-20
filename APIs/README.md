# API 


## gRPC
### Protoc
https://github.com/protocolbuffers/protobuf/releases

### gRPC go plugins
https://grpc.io/docs/languages/go/quickstart/

### gRPC Gateway
https://github.com/grpc-ecosystem/grpc-gateway
```
import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
```

## BUF
https://docs.buf.build/installation