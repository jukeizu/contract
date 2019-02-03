package contract

type Request struct {
	Id        string   `json:"id,omitempty"`
	Source    string   `json:"source,omitempty"`
	Bot       User     `json:"bot,omitempty"`
	Author    User     `json:"author,omitempty"`
	ChannelId string   `json:"channelId,omitempty"`
	ServerId  string   `json:"serverId,omitempty"`
	Servers   []Server `json:"servers,omitempty"`
	Mentions  []User   `json:"mentions,omitempty"`
	Content   string   `json:"content,omitempty"`
}

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Server struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
