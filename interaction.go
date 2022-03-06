package contract

type Interaction struct {
	Identifier string   `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	Source     string   `json:"source,omitempty" yaml:"source,omitempty"`
	MessageId  string   `json:"messageId,omitempty" yaml:"messageId,omitempty"`
	Bot        User     `json:"bot,omitempty" yaml:"bot,omitempty"`
	User       User     `json:"user,omitempty" yaml:"user,omitempty"`
	ChannelId  string   `json:"channelId,omitempty" yaml:"channelId,omitempty"`
	ServerId   string   `json:"serverId,omitempty" yaml:"serverId,omitempty"`
	IsDirect   bool     `json:"isDirect,omitempty" yaml:"isDirect,omitempty"`
	Values     []string `json:"values,omitempty" yaml:"values,omitempty"`
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`
}
