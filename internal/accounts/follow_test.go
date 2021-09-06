package accounts

import "testing"

func TestNewFollowers(t *testing.T) {
	TestInitDB(t)
	// hvturingga fo hvturi
	nf := NewFollow("hvturingga", "hvturi")
	err := nf.New()
	if err != nil {
		return 
	}
}

func TestFollows_Get(t *testing.T) {
	TestInitDB(t)
	// hvturingga fo hvturi
	nf := NewGetFollow("hvturingga")
	nf.Get()
}
