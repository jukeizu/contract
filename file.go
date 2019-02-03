package contract

type File struct {
	Name        string `json:"name,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Bytes       []byte `json:"bytes,omitempty"`
}
