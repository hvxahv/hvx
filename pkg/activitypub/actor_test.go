package activitypub

import "testing"

func TestIsActorExists(t *testing.T) {
	exists, err := IsActorExists("acct:hvturingga@fac47988a686.ngrok.io")
	if err != nil {
		return 
	}
	t.Log(exists)
}

func TestGetHost(t *testing.T) {
	r := GetHost("acct:hvturingga@fac47988a686.ngrok.io")
	t.Log(r)
}