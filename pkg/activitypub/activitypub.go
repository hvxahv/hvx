package activitypub

// SendActivity ... 发送活动的结构体
type Activity struct {
	Data []byte
	EndInbox string
	Method string
	Name string
	UserAddress string
	EndActor string
}

/**
	Data 经过 Marshal 的请求数据
	EndInbox 终点的收件箱，或者发送请求的终点
	Method 请求的方法
	Name 当前用户的用户名
	UserAddress 当前用户的 Activity 地址
	EndActor 终点用户的 Actor 地址； 对方
*/
func NewSendActivity(data []byte, endInbox, method, name, userAddress, endActor string) *Activity {
	nsa := &Activity{
		Data:     data,
		EndInbox: endInbox,
		Method:   method,
		Name:     name,
		UserAddress:  userAddress,
		EndActor: endActor,
	}
	return nsa
}
