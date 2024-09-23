package configs

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type ParserType string

const (
	JsonParser    ParserType = "json"
	YamlParser    ParserType = "yaml"
	DefaultParser ParserType = "default"
)

type TLConfigs struct {
	AddDateTime  bool
	EnableColors bool
	Parser       ParserType
	LogLvl       log_level.LogLevel
}

func NewDefaultTLConfigs() *TLConfigs {
	return &TLConfigs{
		AddDateTime:  false,
		EnableColors: true,
		Parser:       DefaultParser,
		LogLvl: log_level.LogLevel{
			Lvl:         log_level.RetrieveLogLvlIntFromName(log_level.DebugLvlName),
			EnvVariable: "",
		},
	}
}
