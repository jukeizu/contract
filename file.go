package contract

type File struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Bytes       []byte `json:"bytes,omitempty" yaml:"bytes,omitempty"`
}
