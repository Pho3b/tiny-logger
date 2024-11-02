package logs

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/interfaces"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/encoders"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

var (
	packageLoggerInstance interfaces.LoggerInterface
	dateTimeEnabled       bool
	colorsEnabled         bool
	encoder               interfaces.Encoder = encoders.NewDefaultEncoder()
	logLvl                log_level.LogLevel = log_level.LogLevel{
		Lvl:         0,
		EnvVariable: "",
	}
)

func init() {
	packageLoggerInstance = NewLogger()
}

// Debug checks whether the packageLogLvl is sufficiently high and calls the logDebug() method from the package if it is.
func Debug(args ...interface{}) {
	if logLvl.Lvl >= log_level.DebugLvl {
		encoder.LogDebug(packageLoggerInstance, args...)
	}
}

// Info checks whether the packageLogLvl is sufficiently high and calls the logInfo() method from the package if it is.
func Info(args ...interface{}) {
	if logLvl.Lvl >= log_level.InfoLvl {
		logInfo(args...)
	}
}

// Warn checks whether the packageLogLvl is sufficiently high and calls the logWarn() method from the package if it is.
func Warn(args ...interface{}) {
	if logLvl.Lvl >= log_level.WarnLvl {
		logWarn(args...)
	}
}

// Error checks whether the packageLogLvl is sufficiently high and calls the logError() method from the package if it is.
func Error(args ...interface{}) {
	if logLvl.Lvl >= log_level.ErrorLvl {
		logError(args...)
	}
}

// FatalError calls the logFatalError() package method, see its method documentation for more logInfo.
func FatalError(args ...interface{}) {
	logFatalError(args...)
}

// GetLogLvlName returns the package Logger log Level Name.
func GetLogLvlName() log_level.LogLvlName {
	return logLvl.LvlName()
}

// GetLogLvlIntValue returns the package Logger log Level int8 value.
func GetLogLvlIntValue() int8 {
	return logLvl.Lvl
}

// SetLogLvl updates the package Logger log level property if the given logLvlName is valid,
// otherwise it sets the log level to DebugLvlName.
func SetLogLvl(lvlName log_level.LogLvlName) {
	logLvl.Lvl = log_level.RetrieveLogLvlIntFromName(lvlName)
}

// SetLogLvlEnvVariable updates the package Logger log level property attempting to
// retrieve it from the given envVariableName's value.
// If the env variable is not found sets the package log level to DebugLvlName by default.
func SetLogLvlEnvVariable(envVariableName string) {
	logLvl.EnvVariable = envVariableName
	logLvl.Lvl = log_level.RetrieveLogLvlFromEnv(logLvl.EnvVariable)
}
