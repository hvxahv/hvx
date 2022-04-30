package v

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
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
	*port
	err  error
	ctx  context.Context
	conn *grpc.ClientConn
	mux  *runtime.ServeMux
}

type port struct {
	grpc string
	rest string
}

func NewServer(name string, server *grpc.Server, ctx context.Context, mux *runtime.ServeMux) *Server {
	return &Server{
		Server:  server,
		address: viper.GetString(fmt.Sprintf("microservices.%s.address", name)),
		port: &port{
			grpc: viper.GetString(fmt.Sprintf("microservices.%s.gp", name)),
			rest: viper.GetString(fmt.Sprintf("microservices.%s.gwp", name)),
		},
		ctx:  ctx,
		conn: &grpc.ClientConn{},
		mux:  mux,
	}
}

func (c *Cfg) NewServer() *Server {
	return NewServer(
		c.opts.name,
		grpc.NewServer(), context.Background(), runtime.NewServeMux(),
	).ListenerWithEndpoints()
}

func (s *Server) ListenerWithEndpoints() *Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.grpc))
	if err != nil {
		s.err = errors.Wrap(err, "failed to listen")
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", s.address, s.grpc),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		s.err = errors.Wrap(err, "failed to dial server...")
	}

	s.conn = conn
	s.listener = lis
	return s
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

func (s *Server) Run() error {
	if s.err != nil {
		return s.err
	}
	go func() {
		if err := s.Serve(s.listener); err != nil {
			s.err = errors.Wrap(err, "failed to serve")
		}
	}()
	log.Println("grpc server is running on: ", s.grpc)

	gw := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.rest),
		Handler: s.mux,
	}

	log.Println("server gRPC-Gateway is running on: ", s.rest)
	log.Fatalln(gw.ListenAndServe())
	return nil
}
