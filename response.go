package contract

type Response struct {
	Messages []*Message `json:"messages"`
}
