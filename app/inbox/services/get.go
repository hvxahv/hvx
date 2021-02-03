package services

import (
	"fmt"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/pkg/db"
	"log"
)

func GetInboxData(in *pb.Name) []*pb.InboxData {

	db := db.GetMaria()

	var i []*pb.InboxData

	if db.Debug().Table("inbox").Where("name = ?", in.Name).Find(&i).RecordNotFound() {
		log.Println("未查询到用户的 inbox")
	}
	fmt.Println(&i)
	return i
}
