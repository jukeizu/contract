package contract

type Embed struct {
	Url          string         `json:"url"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Timestamp    string         `json:"timestamp"`
	Color        int            `json:"color"`
	Footer       *EmbedFooter   `json:"footer"`
	ImageUrl     string         `json:"imageUrl"`
	ThumbnailUrl string         `json:"thumbnailUrl"`
	VideoUrl     string         `json:"videoUrl"`
	Provider     *EmbedProvider `json:"provider"`
	Author       *EmbedAuthor   `json:"author"`
	Fields       []*EmbedField  `json:"fields"`
}

type EmbedFooter struct {
	Text    string `json:"text"`
	IconUrl string `json:"iconUrl"`
}

type EmbedProvider struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type EmbedAuthor struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
