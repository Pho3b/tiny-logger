package shared

type EncoderType string

const (
	DefaultEncoderType EncoderType = "default"
	JsonEncoderType    EncoderType = "json"
	YamlEncoderType    EncoderType = "yaml"
)

type OutputType int8

const (
	StdOutput    OutputType = 0
	StdErrOutput OutputType = 1
)
