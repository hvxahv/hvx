package services

import (
	"fmt"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/pkg/database"
	"hvxahv/pkg/structs"
	"log"
)

func GetInboxData(in *pb.Name) []*structs.Inbox {

	db := database.GetMaria()

	var i []*structs.Inbox

	if db.Debug().Table("inbox").Where("name = ?", in.Name).Find(&i).RecordNotFound() {
		log.Println("未查询到用户的 inbox")
	}
	fmt.Println(&i)
	return i
}
