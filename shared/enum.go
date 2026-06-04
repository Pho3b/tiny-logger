package shared

type EncoderType string

const (
	DefaultEncoderType EncoderType = "default"
	JsonEncoderType    EncoderType = "json"
	YamlEncoderType    EncoderType = "yaml"
)

type OutputType int8

const (
	StdOutput OutputType = iota
	StdErrOutput
	FileOutput
)

type DateTimeFormat int8

const (
	IT DateTimeFormat = iota
	JP
	US
	UnixTimestamp
)
