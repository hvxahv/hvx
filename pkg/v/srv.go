package v

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type Server struct {
	*grpc.Server
	listener net.Listener
	address  string
	endpoint string
	rest     string
	err      error
	ctx      context.Context
	conn     *grpc.ClientConn
	mux      *runtime.ServeMux
}

func NewServer(name string, server *grpc.Server, ctx context.Context, mux *runtime.ServeMux) *Server {
	return &Server{
		Server:   server,
		address:  GetServiceAddresses(name),
		endpoint: GetServiceEndpoint(name),
		rest:     GetRestEndpoint(name),
		ctx:      ctx,
		conn:     &grpc.ClientConn{},
		mux:      mux,
	}
}

func (c *Cfg) NewServer() *Server {
	return NewServer(
		c.opts.name,
		grpc.NewServer(), context.Background(), runtime.NewServeMux(),
	).ListenerWithEndpoints()
}

func (s *Server) ListenerWithEndpoints() *Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.address))
	if err != nil {
		s.err = errors.Wrap(err, "failed to listen")
	}
	conn, err := grpc.Dial(
		s.endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		s.err = errors.Wrap(err, "Failed to dial server...")
	}
	s.conn = conn
	s.listener = lis
	return s
}

func (s *Server) Run() error {
	go func() {
		if err := s.Serve(s.listener); err != nil {
			s.err = errors.Wrap(err, "failed to serve")
		}
	}()

	log.Println("Server is running on: ", s.endpoint)

	gw := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.rest),
		Handler: s.mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:" + s.rest)
	log.Fatalln(gw.ListenAndServe())
	return nil
}

func (s *Server) GetCtx() context.Context {
	return s.ctx
}

func (s *Server) GetConn() *grpc.ClientConn {
	return s.conn
}

func (s *Server) GetMux() *runtime.ServeMux {
	return s.mux
}
