package channel

type Broadcasts struct {
	Author  string
	Article string
}

func (b Broadcasts) New() {
	panic("implement me")
}

func NewBroadcast(author string, article string) *Broadcasts {
	return &Broadcasts{Author: author, Article: article}
}

type Broadcast interface {
	// New Create broadcast Articles.
	// Synchronize to ipfs return url.
	New()
}