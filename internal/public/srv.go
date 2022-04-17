package public

import (
	"context"
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/v1alpha1/proto/public/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

const serviceName = "public"

type server struct {
	pb.PublicServiceServer
}

func Run() error {
	var port = microservices.NewService(serviceName).GetGRPCPort()
	var restful = microservices.NewService(serviceName).GetHTTPPort()
	s := grpc.NewServer()
	fmt.Println(port)
	pb.RegisterPublicServiceServer(s, &server{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Errorf("failed to listen: %v", err)
	}
	log.Println("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", port))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, err := grpc.Dial(
		microservices.NewService(serviceName).GetGRPCAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()

	err = pb.RegisterPublicServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return errors.Errorf("Failed to register gateway: %v", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", restful),
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0" + fmt.Sprintf(":%d", restful))
	log.Fatalln(gwServer.ListenAndServe())
	return nil
}
