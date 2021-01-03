package bot

import (
	"fmt"
)

func NewAccountNotice(str string)  {
	SendNotice(str)
}


func ServicesRunningNotice(srvname string, port string) {
	str := fmt.Sprintf("%s services is running..., port: %s", srvname, port)
	SendNotice(str)
}