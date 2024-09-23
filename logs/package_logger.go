package logs

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/internal/shared"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/configs"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

// packageLoggerConfigs represents the package logger configs.TLConfigs struct.
var packageLoggerConfigs = configs.NewDefaultTLConfigs()

// Log calls the underlying log() method from the package.
// It always prints the given messages because it does not take the packageLogLvl into account.
func Log(color colors.Color, args ...interface{}) {
	shared.Log(color, args...)
}

// Debug checks whether the packageLogLvl is sufficiently high and calls the logDebug() method from the package if it is.
func Debug(args ...interface{}) {
	if packageLoggerConfigs.LogLvl.Lvl >= log_level.DebugLvl {
		shared.LogDebug(args...)
	}
}

// Info checks whether the packageLogLvl is sufficiently high and calls the logInfo() method from the package if it is.
func Info(args ...interface{}) {
	if packageLoggerConfigs.LogLvl.Lvl >= log_level.InfoLvl {
		shared.LogInfo(args...)
	}
}

// Warn checks whether the packageLogLvl is sufficiently high and calls the logWarn() method from the package if it is.
func Warn(args ...interface{}) {
	if packageLoggerConfigs.LogLvl.Lvl >= log_level.WarnLvl {
		shared.LogWarn(args...)
	}
}

// Error checks whether the packageLogLvl is sufficiently high and calls the logError() method from the package if it is.
func Error(args ...interface{}) {
	if packageLoggerConfigs.LogLvl.Lvl >= log_level.ErrorLvl {
		shared.LogError(args...)
	}
}

// FatalError calls the logFatalError() package method, see its method documentation for more logInfo.
func FatalError(args ...interface{}) {
	shared.LogFatalError(args...)
}

// GetLogLvlName returns the package Logger Log Level Name.
func GetLogLvlName() log_level.LogLvlName {
	return packageLoggerConfigs.LogLvl.LvlName()
}

// GetLogLvlIntValue returns the package Logger Log Level int8 value.
func GetLogLvlIntValue() int8 {
	return packageLoggerConfigs.LogLvl.Lvl
}

// SetLogLvl updates the package Logger log level property if the given logLvlName is valid,
// otherwise it sets the log level to DebugLvlName.
func SetLogLvl(lvlName log_level.LogLvlName) {
	packageLoggerConfigs.LogLvl.Lvl = log_level.RetrieveLogLvlIntFromName(lvlName)
}

// SetLogLvlEnvVariable updates the package Logger log level property attempting to
// retrieve it from the given envVariableName's value.
// If the env variable is not found sets the package log level to DebugLvlName by default.
func SetLogLvlEnvVariable(envVariableName string) {
	packageLoggerConfigs.LogLvl.EnvVariable = envVariableName
	packageLoggerConfigs.LogLvl.Lvl = log_level.RetrieveLogLvlFromEnv(packageLoggerConfigs.LogLvl.EnvVariable)
}
