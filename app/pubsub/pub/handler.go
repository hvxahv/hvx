package main
//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"hvxahv/api/util/v1"
//	"log"
//)
//
//func PubHandler(c *gin.Context) {
//	message := c.Query("message")
//	// 发布客户端
//	conn, err := grpc.Dial("localhost:9010", grpc.WithInsecure())
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//
//	client := v1.NewPubsubServiceClient(conn)
//
//	str, err := client.Publish(context.Background(), &v1.PubSubMessage{Message: message})
//
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		fmt.Println("Publish 返回的消息: ", str.GetMessage())
//	}
//}
