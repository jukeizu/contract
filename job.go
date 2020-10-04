package contract

type Job struct {
	Id          string   `json:"id,omitempty" yaml:"id,omitempty"`
	UserId      string   `json:"userId,omitempty" yaml:"userId,omitempty"`
	InstanceId  string   `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`
	Source      string   `json:"source,omitempty" yaml:"source,omitempty"`
	Content     string   `json:"content,omitempty" yaml:"content,omitempty"`
	Endpoint    string   `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Destination string   `json:"destination,omitempty" yaml:"destination,omitempty"`
	Schedule    Schedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`
	Enabled     bool     `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Created     int32    `json:"created,omitempty" yaml:"created,omitempty"`
}

type Schedule struct {
	Minute     string `json:"minute,omitempty" yaml:"minute,omitempty"`
	Hour       string `json:"hour,omitempty" yaml:"hour,omitempty"`
	DayOfMonth string `json:"dayOfMonth,omitempty" yaml:"dayOfMonth,omitempty"`
	Month      string `json:"month,omitempty" yaml:"month,omitempty"`
	DayOfWeek  string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`
	Year       string `json:"year,omitempty" yaml:"year,omitempty"`
}
