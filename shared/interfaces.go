package shared

import (
	"os"

	"github.com/Pho3b/tiny-logger/logs/colors"
	"github.com/Pho3b/tiny-logger/logs/log_level"
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
	GetLogFile() *os.File
	GetDateTimeFormat() DateTimeFormat
}

type EncoderInterface interface {
	Log(logger LoggerConfigsInterface, lvl log_level.LogLvlName, outType OutputType, args ...any)
	Color(lConfigs LoggerConfigsInterface, color colors.Color, args ...any)
	GetType() EncoderType
}
