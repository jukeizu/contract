package contract

type Job struct {
	Id          string `json:"id,omitempty"`
	UserId      string `json:"userId,omitempty"`
	Source      string `json:"source,omitempty"`
	Content     string `json:"content,omitempty"`
	Endpoint    string `json:"endpoint,omitempty"`
	Destination string `json:"destination,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Created     int32  `json:"created,omitempty"`
}
