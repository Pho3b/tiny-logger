package shared

import (
	"os"

	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/log_level"
)

type LoggerInterface interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	FatalError(args ...any)
}

type LoggerConfigsInterface interface {
	GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool)
	GetColorsEnabled() bool
	GetShowLogLevel() bool
	GetLogLvlName() log_level.LogLvlName
	GetLogLvlIntValue() int8
	GetEncoderType() EncoderType
	IsFileLogEnabled() bool
}

type EncoderInterface interface {
	LogDebug(lConfigs LoggerConfigsInterface, args ...any)
	LogInfo(lConfigs LoggerConfigsInterface, args ...any)
	LogWarn(lConfigs LoggerConfigsInterface, args ...any)
	LogError(lConfigs LoggerConfigsInterface, args ...any)
	LogFatalError(lConfigs LoggerConfigsInterface, args ...any)
	Color(lConfigs LoggerConfigsInterface, color colors.Color, args ...any)
	GetType() EncoderType
	GetOutFile() *os.File
	SetOutFile(outFile *os.File)
}
