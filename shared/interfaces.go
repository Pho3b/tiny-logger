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
	LogDebug(lConfigs LoggerConfigsInterface, args ...interface{})
	LogInfo(lConfigs LoggerConfigsInterface, args ...interface{})
	LogWarn(lConfigs LoggerConfigsInterface, args ...interface{})
	LogError(lConfigs LoggerConfigsInterface, args ...interface{})
	LogFatalError(lConfigs LoggerConfigsInterface, args ...interface{})
	Color(lConfigs LoggerConfigsInterface, color colors.Color, args ...interface{})
	GetType() EncoderType
}
