package contract

type Message struct {
	Content          string      `json:"content,omitempty" yaml:"content,omitempty"`
	Embed            *Embed      `json:"embed,omitempty" yaml:"embed,omitempty"`
	Reactions        []string    `json:"reactions,omitempty" yaml:"reactions,omitempty"`
	Files            []*File     `json:"files,omitempty" yaml:"files,omitempty"`
	Compontents      *Components `json:"components,omitempty" yaml:"components,omitempty"`
	Tts              bool        `json:"tts,omitempty" yaml:"tts,omitempty"`
	IsPrivateMessage bool        `json:"isPrivateMessage,omitempty" yaml:"isPrivateMessage,omitempty"`
	IsRedirect       bool        `json:"isRedirect,omitempty" yaml:"isRedirect,omitempty"`
	EditMessageId    string      `json:"editMessageId,omitempty" yaml:"editMessageId,omitempty"`
}
