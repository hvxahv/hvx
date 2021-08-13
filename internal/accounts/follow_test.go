package accounts

import "testing"

func TestNewFollowers(t *testing.T) {
	TestInitDB(t)
	// hvturingga fo hvturi
	nf := NewFollowers("hvturingga", "hvturi")
	err := nf.New()
	if err != nil {
		return 
	}
}
