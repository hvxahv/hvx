package activitypub

// Object Describes an object of any kind.
// The Object type serves as the base type for most of the other kinds of objects defined in the Activity Vocabulary,
// including other Core types such as Activity,
// IntransitiveActivity, Collection and OrderedCollection.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-object
//type Object struct {
//	Context string `json:"@context"`
//	Type    string `json:"type"`
//	Id      string `json:"id"`
//	Name    string `json:"name"`
//}

// Article Represents any kind of multi-paragraph written work.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-article
type Article struct {
	Context      string `json:"@context"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	AttributedTo string `json:"attributedTo"`
}

// Document Represents a document of any kind.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-document
type Document struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

// Audio Represents an audio document of any kind.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-audio
type Audio struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Url     struct {
		Type      string `json:"type"`
		Href      string `json:"href"`
		MediaType string `json:"mediaType"`
	} `json:"url"`
}

// Image An image document of any kind.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-image
type Image struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Url     []struct {
		Type      string `json:"type"`
		Href      string `json:"href"`
		MediaType string `json:"mediaType"`
	} `json:"url"`
}

// Note Represents a short written work typically less than a single paragraph in length.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-note
type Note struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Place Represents a logical or physical location. See 5.3 Representing Places for additional information.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-place
type Place struct {
	Context   string  `json:"@context"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    int     `json:"radius"`
	Units     string  `json:"units"`
}

// Video Represents a video document of any kind.
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-video
type Video struct {
	Context  string `json:"@context"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Duration string `json:"duration"`
}
