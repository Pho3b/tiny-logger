package shared

type JsonLog struct {
	Level    string         `json:"level,omitempty"`
	Date     string         `json:"date,omitempty"`
	Time     string         `json:"time,omitempty"`
	DateTime string         `json:"datetime,omitempty"`
	Message  string         `json:"msg"`
	Extras   map[string]any `json:"extras,omitempty"`
}
