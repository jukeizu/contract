package contract

type Embed struct {
	Url          string        `json:"url,omitempty"`
	Title        string        `json:"title,omitempty"`
	Description  string        `json:"description,omitempty"`
	Timestamp    string        `json:"timestamp,omitempty"`
	Color        int           `json:"color,omitempty"`
	Footer       *EmbedFooter  `json:"footer,omitempty"`
	ImageUrl     string        `json:"imageUrl,omitempty"`
	ThumbnailUrl string        `json:"thumbnailUrl,omitempty"`
	VideoUrl     string        `json:"videoUrl,omitempty"`
	Author       *EmbedAuthor  `json:"author,omitempty"`
	Fields       []*EmbedField `json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text    string `json:"text,omitempty"`
	IconUrl string `json:"iconUrl,omitempty"`
}

type EmbedAuthor struct {
	Url  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}
