package microsvc

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
	"sync"
)

type server struct {
	Name string
	*grpc.Server
	Ctx      context.Context
	Cancel   context.CancelFunc
	Listener net.Listener
	wg       sync.WaitGroup
	mutex    sync.RWMutex
	stop     chan struct{}
	done     chan struct{}
	Err      error
	Mux      *runtime.ServeMux
	Conn     *grpc.ClientConn
	*URLs
}

type URLs struct {
	host string
	http string
	grpc string
}

func (c *Cfg) ListenerWithEndpoints() *server {
	var (
		host     = viper.GetString(fmt.Sprintf("microsvcs.%s.hostname", c.opts.name))
		grpcPort = viper.GetString(fmt.Sprintf("microsvcs.%s.ports.grpc", c.opts.name))
		httpPort = viper.GetString(fmt.Sprintf("microsvcs.%s.ports.http", c.opts.name))
		ctx      = context.Background()
		mux      = runtime.NewServeMux()
		err      error
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		err = errors.Wrap(err, "failed to listen")
	}

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", host, grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		err = errors.Wrap(err, "failed to dial server...")
	}

	return &server{
		Name:     c.opts.name,
		Server:   grpc.NewServer(),
		Ctx:      ctx,
		Cancel:   nil,
		Listener: listener,
		wg:       sync.WaitGroup{},
		mutex:    sync.RWMutex{},
		stop:     nil,
		done:     nil,
		Err:      err,
		Mux:      mux,
		Conn:     conn,
		URLs: &URLs{
			host: host,
			http: httpPort,
			grpc: grpcPort,
		},
	}
}

func (s *server) Run() error {
	// Service registration...
	//r, err := consul.New(s.Name, s.host, s.http, []string{"http", "grpc", "microservices"})
	//if err != nil {
	//	return err
	//}
	//
	//if err := r.Register(); err != nil {
	//	return err
	//}

	s.wg.Add(1)
	go func() {
		if err := s.Serve(s.Listener); err != nil {
			s.Err = errors.Wrap(err, "failed to serve")
			s.wg.Done()
		}
	}()

	log.Println("grpc server is running on: ", s.grpc)

	gw := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.http),
		Handler: s.Mux,
	}

	log.Println("server gRPC-Gateway is running on: ", s.http)

	if err := gw.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
