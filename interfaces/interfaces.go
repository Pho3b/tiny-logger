package interfaces

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type Encoder interface {
	LogDebug(logger LoggerInterface, args ...interface{})
}

type LoggerInterface interface {
	Log(color colors.Color, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	FatalError(args ...interface{})
	GetDateTimeEnabled() bool
	GetColorsEnabled() bool
	GetLogLvlName() log_level.LogLvlName
	GetLogLvlIntValue() int8
	SetLogLvl(logLvlName log_level.LogLvlName) LoggerInterface
	SetEnableColors(enable bool) LoggerInterface
	SetAddDateTime(addDateTime bool) LoggerInterface
	SetLogLvlEnvVariable(envVariableName string) LoggerInterface
}
