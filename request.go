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
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Discriminator string `json:"discriminator"`
}

type Server struct {
	Id              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	OwnerId         string `json:"ownerId,omitempty"`
	Description     string `json:"description,omitempty"`
	UserCount       int32  `json:"userCount,omitempty"`
	IconUrl         string `json:"iconUrl,omitempty"`
	SystemChannelId string `json:"systemChannelId,omitempty"`
}

// Server returns the full server for the ServerId
func (r Request) Server() Server {
	for _, server := range r.Servers {
		if server.Id == r.ServerId {
			return server
		}
	}

	return Server{Id: r.ServerId}
}
