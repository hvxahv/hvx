package api_client

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)

// Conn .. gRPC 客户端连接服务端的方法封装，它接收两个参数，服务的端口号和服务名并返回 conn
func Conn(port, name string) (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("localhost:%s",port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("连接到 %s 服务失败: %v", err, name)
	}
	return conn, err
}
