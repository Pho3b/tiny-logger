package configs

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/internal/services"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/encoders"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type TLConfigs struct {
	AddDateTime  bool
	EnableColors bool
	Encoder      encoders.Encoder
	LogLvl       log_level.LogLevel
}

func NewDefaultTLConfigs() *TLConfigs {
	return &TLConfigs{
		AddDateTime:  false,
		EnableColors: true,
		Encoder: encoders.NewDefaultEncoder(
			services.Wrapper{
				DateTimePrinter: services.DateTimePrinter{},
				ColorsPrinter:   services.ColorsPrinter{},
			},
		),
		LogLvl: log_level.LogLevel{
			Lvl:         log_level.RetrieveLogLvlIntFromName(log_level.DebugLvlName),
			EnvVariable: "",
		},
	}
}
