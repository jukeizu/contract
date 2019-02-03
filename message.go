package contract

type Message struct {
	Content          string  `json:"content,omitempty"`
	Embed            *Embed  `json:"embed,omitempty"`
	Files            []*File `json:"files,omitempty"`
	Tts              bool    `json:"tts,omitempty"`
	IsPrivateMessage bool    `json:"isPrivateMessage,omitempty"`
	IsRedirect       bool    `json:"isRedirect,omitempty"`
}
