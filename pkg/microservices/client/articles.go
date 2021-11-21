package client

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/articles/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

func Articles() (pb.ArticlesClient, error) {
	name := "articles"

	host := viper.GetString(fmt.Sprintf("microservices.%s.host", name))
	port := viper.GetString(fmt.Sprintf("microservices.%s.port", name))
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to %s service: %v", name , err)
		return nil, err
	}

	defer conn.Close()
	cli := pb.NewArticlesClient(conn)

	return cli, nil
}