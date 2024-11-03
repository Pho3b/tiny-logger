package log_level

import (
	"os"
	"strings"
)

// LogLvlName is the Enum representing the possible Log Levels.
type LogLvlName string

// DefaultEnvLogLvlVar is the default ENV variable that any new logger will try to
// retrieve the loge level from by default.
const DefaultEnvLogLvlVar = "TINY_LOGGER_LVL"

const (
	ErrorLvlName LogLvlName = "ERROR"
	WarnLvlName  LogLvlName = "WARN"
	InfoLvlName  LogLvlName = "INFO"
	DebugLvlName LogLvlName = "DEBUG"
)

// Log Level INT8 Constants
const (
	ErrorLvl = int8(0)
	WarnLvl  = int8(1)
	InfoLvl  = int8(2)
	DebugLvl = int8(3)
)

// LogLvlIntToName represents the log level INT to STRING map
var LogLvlIntToName = map[int8]LogLvlName{
	ErrorLvl: ErrorLvlName,
	WarnLvl:  WarnLvlName,
	InfoLvl:  InfoLvlName,
	DebugLvl: DebugLvlName,
}

// LogLvlNameToInt represents the log level STRING to INT map
var LogLvlNameToInt = map[LogLvlName]int8{
	ErrorLvlName: ErrorLvl,
	WarnLvlName:  WarnLvl,
	InfoLvlName:  InfoLvl,
	DebugLvlName: DebugLvl,
}

type LogLevel struct {
	Lvl         int8
	EnvVariable string
}

func (l *LogLevel) LvlName() LogLvlName {
	return LogLvlIntToName[l.Lvl]
}

func (l *LogLevel) LvlIntValue() int8 {
	return l.Lvl
}

// RetrieveLogLvlFromEnv attempts to retrieve the given 'logLvlEnvVariable' value from the ENV variables.
// If the given variable is not found, 'DebugLvl' is returned by default.
func RetrieveLogLvlFromEnv(logLvlEnvVariable string) int8 {
	return RetrieveLogLvlIntFromName(
		LogLvlName(
			strings.ToUpper(os.Getenv(logLvlEnvVariable)),
		),
	)
}

// RetrieveLogLvlIntFromName given a logLvlName returns its int8 constant value.
// If the given logLvlName is not valid, DebugLvl is returned by default.
func RetrieveLogLvlIntFromName(logLvlName LogLvlName) int8 {
	if _, found := LogLvlNameToInt[logLvlName]; !found {
		logLvlName = DebugLvlName
	}

	return LogLvlNameToInt[logLvlName]
}
