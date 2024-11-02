package interfaces

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type Encoder interface {
	LogDebug(logger LoggerInterface, args ...interface{})
	LogInfo(logger LoggerInterface, args ...interface{})
	LogWarn(logger LoggerInterface, args ...interface{})
	LogError(logger LoggerInterface, args ...interface{})
	LogFatalError(logger LoggerInterface, args ...interface{})
}

type LoggerInterface interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	FatalError(args ...interface{})
	GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool)
	GetColorsEnabled() bool
	GetLogLvlName() log_level.LogLvlName
	GetLogLvlIntValue() int8
	SetLogLvl(logLvlName log_level.LogLvlName) LoggerInterface
	EnableColors(enable bool) LoggerInterface
	AddDateTime(addDateTime bool) LoggerInterface
	AddDate(addDate bool) LoggerInterface
	AddTime(addTime bool) LoggerInterface
	SetLogLvlEnvVariable(envVariableName string) LoggerInterface
}
