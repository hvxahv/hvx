package microservice

import pb "hvxahv/api/hvxahv/v1alpha1"

type microservice struct {
	Name string
	Addr string
	Port string
}

type server struct {
	pb.AccountsServer
	pb.ArticlesServer
}


// Server hvxahv server interface for all microservices
type Server interface {
	AccountServer() error
	ArticlesServer() error
}

// NewServer It accepts the name and port number of the microservice and returns the Server interface.
func NewServer(name string, port string) Server {
	return &microservice{Name: name, Port: port}
}
