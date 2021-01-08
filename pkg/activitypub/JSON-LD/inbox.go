package JSON_LD

import "time"

type Inbox struct {
	Context string `json:"@context"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Object  struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Actor  string `json:"actor"`
		Object string `json:"object"`
	} `json:"object"`
}

// Messages 私信的结构体
type Messages struct {
	Context   []interface{} `json:"@context"`
	ID        string        `json:"id"`
	Type      string        `json:"type"`
	Actor     string        `json:"actor"`
	Published time.Time     `json:"published"`
	To        []string      `json:"to"`
	Cc        []interface{} `json:"cc"`
	Object    struct {
		ID               string        `json:"id"`
		Type             string        `json:"type"`
		Summary          interface{}   `json:"summary"`
		InReplyTo        interface{}   `json:"inReplyTo"`
		Published        time.Time     `json:"published"`
		URL              string        `json:"url"`
		AttributedTo     string        `json:"attributedTo"`
		To               []string      `json:"to"`
		Cc               []interface{} `json:"cc"`
		Sensitive        bool          `json:"sensitive"`
		AtomURI          string        `json:"atomUri"`
		InReplyToAtomURI interface{}   `json:"inReplyToAtomUri"`
		Conversation     string        `json:"conversation"`
		Content          string        `json:"content"`
		ContentMap       struct {
			ZhCn string `json:"zhCn"`
		} `json:"contentMap"`
		Attachment []interface{} `json:"attachment"`
		Tag        []struct {
			Type string `json:"type"`
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"tag"`
		Replies struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			First struct {
				Type   string        `json:"type"`
				Next   string        `json:"next"`
				PartOf string        `json:"partOf"`
				Items  []interface{} `json:"items"`
			} `json:"first"`
		} `json:"replies"`
	} `json:"object"`
}