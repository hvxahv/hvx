package main
//
//import (
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"hvxahv/api/util/v1alpha1"
//	"io"
//	"log"
//)
//
//func main()  {
//	// 订阅客户端
//	conn, err := grpc.Dial("localhost:9010", grpc.WithInsecure())
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//
//	c := v1alpha1.NewPubsubServiceClient(conn)
//
//	stream, err := c.Subscribe(context.Background(), &v1alpha1.PubSubMessage{Message: "nmsl"})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for {
//		reply, err := stream.Recv()
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			log.Fatal(err)
//		}
//		log.Println(reply.GetMessage())
//	}
//}
