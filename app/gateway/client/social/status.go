/*
这个文件中包含 添加 删除 修改 状态的 Handler
连接到 Status Services 并发送前端接收的数据
*/

package social

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hvxahv/api/kernel/v1"
	"log"
)

// CreateStatusHandler 创建状态 handler
func CreateStatusClient(author string) {
	conn, err := grpc.Dial(viper.GetString("port.status"), grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Status app: %v", err)
	}
	defer conn.Close()
	cli := pb.NewStatusClient(conn)
	r, err := cli.NewStatus(context.Background(), &pb.StatusData{
		Author: author,
		Status: "这是我的第一条 Status",
	})
	if err != nil {
		log.Printf("新建状态错误，发送消息给 Status 服务端失败: %v", err)
	}
	log.Println("创建状态返回的数据", r.Reply)
}

// UpdateStatusListHandler 更新状态 Handler
func UpdateStatusClient(author string) {
	addr := fmt.Sprintf("localhost:%s", viper.GetString("port.status"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Status app: %v", err)
	}
	defer conn.Close()
	cli := pb.NewStatusClient(conn)
	r, err := cli.UpdateStatus(context.Background(), &pb.StatusData{
		Id: "123123",
		Author: author,
		Status: "这是我的第一条 Status 的修改内容（开始更新状态）",
	})
	if err != nil {
		log.Printf("更新状态错误，发送消息给 Status 服务端失败: %v", err)
	}
	log.Println("更新状态返回的数据", r.Reply)
}

// DeleteStatusHandler 删除状态 Handler 通过 ID
func DeleteStatusClient(id string) {
	addr := fmt.Sprintf("localhost:%s", viper.GetString("port.status"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Status app: %v", err)
	}
	defer conn.Close()
	cli := pb.NewStatusClient(conn)
	r, err := cli.DeleteStatus(context.Background(), &pb.DeleteStatusByID{
		Id: id,
	})
	if err != nil {
		log.Printf("删除状态错误，发送消息给 Status 服务端失败: %v", err)
	}
	log.Println("删除状态返回的数据", r.Reply)
}
