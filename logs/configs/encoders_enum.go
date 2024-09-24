package configs

type EncoderType string

const (
	JsonEncoderType    EncoderType = "json"
	YamlEncoderType    EncoderType = "yaml"
	DefaultEncoderType EncoderType = "default"
)
