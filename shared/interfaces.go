package shared

import (
	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/log_level"
)

type LoggerInterface interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	FatalError(args ...interface{})
}

type LoggerConfigsInterface interface {
	GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool)
	GetColorsEnabled() bool
	GetShowLogLevel() bool
	GetLogLvlName() log_level.LogLvlName
	GetLogLvlIntValue() int8
	GetEncoderType() EncoderType
}

type EncoderInterface interface {
	LogDebug(logger LoggerConfigsInterface, args ...interface{})
	LogInfo(logger LoggerConfigsInterface, args ...interface{})
	LogWarn(logger LoggerConfigsInterface, args ...interface{})
	LogError(logger LoggerConfigsInterface, args ...interface{})
	LogFatalError(logger LoggerConfigsInterface, args ...interface{})
	GetType() EncoderType
}

type ColorsInterface interface {
	Color(color colors.Color, args ...interface{})
}
