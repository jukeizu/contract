package contract

type File struct {
	Name        string `json:"name"`
	ContentType string `json:"contentType"`
	Bytes       []byte `json:"bytes"`
}
