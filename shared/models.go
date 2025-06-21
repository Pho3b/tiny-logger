package shared

// JsonLog represents the structure of a JSON log and can be used to unmarshal JSON logs.
type JsonLog struct {
	Level    string         `json:"level,omitempty"`
	Date     string         `json:"date,omitempty"`
	Time     string         `json:"time,omitempty"`
	DateTime string         `json:"datetime,omitempty"`
	Message  string         `json:"msg"`
	Extras   map[string]any `json:"extras,omitempty"`
}

// YamlLog represents the structure of a YAML log and can be used to unmarshal YAML logs.
type YamlLog struct {
	Level    string         `yaml:"level,omitempty"`
	Date     string         `yaml:"date,omitempty"`
	Time     string         `yaml:"time,omitempty"`
	DateTime string         `yaml:"datetime,omitempty"`
	Message  string         `yaml:"msg"`
	Extras   map[string]any `yaml:"extras,omitempty"`
}
