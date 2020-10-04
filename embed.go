package contract

type Embed struct {
	Url          string        `json:"url,omitempty" yaml:"url,omitempty"`
	Title        string        `json:"title,omitempty" yaml:"title,omitempty"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	Timestamp    string        `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
	Color        int           `json:"color,omitempty" yaml:"color,omitempty"`
	Footer       *EmbedFooter  `json:"footer,omitempty" yaml:"footer,omitempty"`
	ImageUrl     string        `json:"imageUrl,omitempty" yaml:"imageUrl,omitempty"`
	ThumbnailUrl string        `json:"thumbnailUrl,omitempty" yaml:"thumbnailUrl,omitempty"`
	VideoUrl     string        `json:"videoUrl,omitempty" yaml:"videoUrl,omitempty"`
	Author       *EmbedAuthor  `json:"author,omitempty" yaml:"author,omitempty"`
	Fields       []*EmbedField `json:"fields,omitempty" yaml:"fields,omitempty"`
}

type EmbedFooter struct {
	Text    string `json:"text,omitempty" yaml:"text,omitempty"`
	IconUrl string `json:"iconUrl,omitempty" yaml:"iconUrl,omitempty"`
}

type EmbedAuthor struct {
	Url  string `json:"url,omitempty" yaml:"url,omitempty"`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty" yaml:"name,omitempty"`
	Value  string `json:"value,omitempty" yaml:"value,omitempty"`
	Inline bool   `json:"inline,omitempty" yaml:"inline,omitempty"`
}
