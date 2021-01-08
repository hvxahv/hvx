package JSON_LD

type WebFinger struct {
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
		Type string `json:"type"`
	} `json:"links"`
	Subject string `json:"subject"`
}
