package contract

type Message struct {
	Content          string  `json:"content"`
	Embed            *Embed  `json:"embed"`
	Files            []*File `json:"files"`
	Tts              bool    `json:"tts"`
	IsPrivateMessage bool    `json:"isPrivateMessage"`
	IsRedirect       bool    `json:"isRedirect"`
}
