package contract

type Response struct {
	Messages  []*Message  `json:"messages,omitempty"`
	Reactions []*Reaction `json:"reactions,omitempty"`
}

func StringResponse(content string) *Response {
	message := Message{
		Content: content,
	}

	return &Response{Messages: []*Message{&message}}
}
