package internal

type Attachments struct {
	Attachment []struct {
		Type      string      `json:"type"`
		MediaType string      `json:"mediaType"`
		Url       string      `json:"url"`
		Name      interface{} `json:"name"`
		Blurhash  string      `json:"blurhash"`
		Width     int         `json:"width"`
		Height    int         `json:"height"`
	} `json:"attachment"`
}
