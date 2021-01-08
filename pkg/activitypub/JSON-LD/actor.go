package JSON_LD

type Actor struct {
	Context           []string `json:"@context"`
	ID                string   `json:"id"`
	Type              string   `json:"type"`
	PreferredUsername string   `json:"preferredUsername"`
	Inbox             string   `json:"inbox"`
	PublicKey         struct {
		ID           	string `json:"id"`
		Owner        	string `json:"owner"`
		PublicKeyPem 	string `json:"publicKeyPem"`
	} `json:"publicKey"`
}

func (a *Actor) NewActor() {

}
