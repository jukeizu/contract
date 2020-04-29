package contract

type Job struct {
	Id          string   `json:"id,omitempty"`
	UserId      string   `json:"userId,omitempty"`
	InstanceId  string   `json:"instanceId,omitempty"`
	Source      string   `json:"source,omitempty"`
	Content     string   `json:"content,omitempty"`
	Endpoint    string   `json:"endpoint,omitempty"`
	Destination string   `json:"destination,omitempty"`
	Schedule    Schedule `json:"schedule,omitempty"`
	Enabled     bool     `json:"enabled,omitempty"`
	Created     int32    `json:"created,omitempty"`
}

type Schedule struct {
	Minute     string `json:"minute,omitempty"`
	Hour       string `json:"hour,omitempty"`
	DayOfMonth string `json:"dayOfMonth,omitempty"`
	Month      string `json:"month,omitempty"`
	DayOfWeek  string `json:"dayOfWeek,omitempty"`
	Year       string `json:"year,omitempty"`
}
