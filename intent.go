package contract

type Intent struct {
	Id       string    `json:"id,omitempty" yaml:"id,omitempty"`
	ServerID string    `json:"serverId,omitempty" yaml:"serverId,omitempty"`
	Name     string    `json:"name,omitempty" yaml:"name,omitempty"`
	Regex    string    `json:"regex,omitempty" yaml:"regex,omitempty"`
	Mention  bool      `json:"mention,omitempty" yaml:"mention,omitempty"`
	Endpoint string    `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Help     string    `json:"help,omitempty" yaml:"help,omitempty"`
	Enabled  bool      `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Type     string    `json:"type,omitempty" yaml:"type,omitempty"`
	Response *Response `json:"response,omitempty" yaml:"response,omitempty"`
}
