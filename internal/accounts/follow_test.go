package accounts

import "testing"

func TestNewFollowers(t *testing.T) {
	TestInitDB(t)
	// hvturingga fo hvturi
	nf := NewFollows("hvturingga", "hvturi")
	err := nf.New()
	if err != nil {
		return
	}

	//nf := NewFollows("hvturi", "hvturingga")
	//err := nf.New()
	//if err != nil {
	//	return
	//}
}

func TestFollows_Get(t *testing.T) {
	TestInitDB(t)
	// hvturingga fo hvturi
	//nf := ("hvturingga")
	//nf.Get()
}
