package inbox

// Accept ...
// Actor 远程用户
// Name 当前登录用户名
// RequestId 请求 ID
type Accept struct {
	Actor    	string
	Name		string
	RequestId 	string
}

// NewAccept ...
func NewAccept(actor, name, reqId string) *Accept {
	accept := &Accept{
		Actor: actor,
		Name: name,
		RequestId: reqId,
	}
	return accept
}
