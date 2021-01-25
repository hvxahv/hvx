package main
//
//import (
//	"fmt"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/reflection"
//	v1 "hvxahv/api/util/v1"
//
//	"hvxahv/pkg/pubsub"
//	"log"
//	"net"
//	"strings"
//	"time"
//)
//
//type PubsubService struct {
//	pub *pubsub.Publisher
//	v1.PubsubServiceServer
//}
//
//func NewPubsubService() *PubsubService {
//	return &PubsubService{
//		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
//	}
//}
//
//func (p *PubsubService) Publish(ctx context.Context, arg *v1.PubSubMessage) (*v1.PubSubMessage, error) {
//	str := arg.GetMessage()
//	p.pub.Publish(str)
//	log.Println("接受到的推送者的消息",str)
//	return &v1.PubSubMessage{Message: str + "还给你"}, nil
//}
//
//func (p *PubsubService) Subscribe(arg *v1.PubSubMessage, stream v1.PubsubService_SubscribeServer) error {
//	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
//		if key, ok := v.(string); ok {
//			// 接受订阅者发过来的字符串，如果带这个字符串，订阅者就接受这个消息
//			if strings.Contains(key, arg.GetMessage()) {
//				fmt.Println("订户要这个消息")
//				return true
//			}
//		}
//		fmt.Println("订户不要这个消息")
//		return false
//	})
//
//	for v := range ch {
//		log.Println("发送给订阅者的消息", v)
//		if err := stream.Send(&v1.PubSubMessage{Message: v.(string)}); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//const (
//	spp = ":9010"
//)
//func main() {
//	// Pubsub 服务端
//	lis, err := net.Listen("tcp", spp)
//	if err != nil {
//		fmt.Printf("Failed to Listen: %v", err)
//		return
//	} else {
//		log.Println("PubSub gRPC service is running", spp)
//	}
//
//	s := grpc.NewServer()
//	v1.RegisterPubsubServiceServer(s, NewPubsubService())
//	reflection.Register(s)
//
//	if err := s.Serve(lis); err != nil {
//		fmt.Printf("PubSub gRPC service failed to start: %v", err)
//		return
//	}
//
//
//}