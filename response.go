package contract

type Response struct {
	Messages []*Message `json:"messages,omitempty"`
}

func StringResponse(content string) *Response {
	message := Message{
		Content: content,
	}

	return &Response{Messages: []*Message{&message}}
}
