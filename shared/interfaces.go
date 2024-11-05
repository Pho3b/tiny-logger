package shared

import (
	"github.com/pho3b/tiny-logger/logs/log_level"
)

type EncoderInterface interface {
	LogDebug(logger LoggerConfigsInterface, args ...interface{})
	LogInfo(logger LoggerConfigsInterface, args ...interface{})
	LogWarn(logger LoggerConfigsInterface, args ...interface{})
	LogError(logger LoggerConfigsInterface, args ...interface{})
	LogFatalError(logger LoggerConfigsInterface, args ...interface{})
	GetType() EncoderType
}

type LoggerInterface interface {
	LoggerConfigsInterface
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
	SetEncoder(encoderType EncoderType) LoggerInterface
}

type LoggerConfigsInterface interface {
	GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool)
	GetColorsEnabled() bool
	GetLogLvlName() log_level.LogLvlName
	GetLogLvlIntValue() int8
	GetEncoderType() EncoderType
}
