package services

import (
	"encoding/json"
	"fmt"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/pkg/database"
	"hvxahv/pkg/structs"
	"log"
)

// ReceiveInbox 接收 Inbox 数据
func ReceiveInbox(in *pb.InboxData) string {
	i := &structs.Inbox{
		Actor: in.Actor,
		RequestId: in.RequestId,
		EventType: in.EventType,
		Name: in.Name,
	}
	k := fmt.Sprintf("%s-inbox", in.Name)

	// 存储到 Redis 缓存
	rdb := database.GetRDB()
	v, _ := json.Marshal(&i)
	if _, err := rdb.Do("SETNX", k, v); err != nil {
		log.Printf("Actor 持久化到数据库失败: %s", err)
	}
	// 在协程中将数据保存到 MariaDB 数据库
	go func() {
		db := database.GetMaria()
		db.AutoMigrate(&structs.Inbox{})
		if err := db.Debug().Table("inbox").Create(&i).Error; err != nil {
			log.Println("保存 inbox 消息失败", err)
		}
	}()
	return "ok"
}
