package models

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
