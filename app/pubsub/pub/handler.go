package main
//
//import (
//	"fmt"
//	"github.com/mw-gonic/mw"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"hvxahv/api/util/v1alpha1"
//	"log"
//)
//
//func PubHandler(c *mw.Context) {
//	message := c.Query("message")
//	// 发布客户端
//	conn, err := grpc.Dial("localhost:9010", grpc.WithInsecure())
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//
//	client := v1alpha1.NewPubsubServiceClient(conn)
//
//	str, err := client.Publish(context.Background(), &v1alpha1.PubSubMessage{Message: message})
//
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		fmt.Println("Publish 返回的消息: ", str.GetMessage())
//	}
//}
