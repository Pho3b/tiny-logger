package shared

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type EncoderInterface interface {
	LogDebug(logger LoggerConfigsInterface, args ...interface{})
	LogInfo(logger LoggerConfigsInterface, args ...interface{})
	LogWarn(logger LoggerConfigsInterface, args ...interface{})
	LogError(logger LoggerConfigsInterface, args ...interface{})
	LogFatalError(logger LoggerConfigsInterface, args ...interface{})
}

type LoggerInterface interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	FatalError(args ...interface{})
	SetLogLvl(logLvlName log_level.LogLvlName) LoggerInterface
	EnableColors(enable bool) LoggerInterface
	AddDateTime(addDateTime bool) LoggerInterface
	AddDate(addDate bool) LoggerInterface
	AddTime(addTime bool) LoggerInterface
	SetLogLvlEnvVariable(envVariableName string) LoggerInterface
	SetEncoder(encoderType EncoderType)
	GetCurrentEncoder(encoder EncoderInterface)
}

type LoggerConfigsInterface interface {
	GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool)
	GetColorsEnabled() bool
	GetLogLvlName() log_level.LogLvlName
	GetLogLvlIntValue() int8
}