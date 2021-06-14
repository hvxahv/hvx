package microservice

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)

// Client all microservice client interfaces.
type Client interface {
	NewConn() (*grpc.ClientConn, error)
}

// NewClient He accepts the name and port number of the microservice,
// and finally gets the address of the grpc client
// and returns to the Client interface.
func NewClient(name string, port string) Client {
	addr := fmt.Sprintf("localhost:%s", port)
	return &microservice{Name: name, Addr: addr}
}

// NewConn Create a new client connection and return grpc.ClientConn.
func (ms *microservice) NewConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(ms.Addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to %s service: %v",  ms.Name, err)
	}
	return conn, nil
}
