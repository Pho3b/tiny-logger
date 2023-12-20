package logs

import "gitlab.com/docebo/libraries/go/tiny-logger/colors"

// packageLogLvl represents the package logger log level.
var packageLogLvl = LogLevel{
	lvl:         retrieveLogLvlIntFromName(DebugLvlName),
	envVariable: "",
}

// Log calls the underlying log() method from the package.
// It always prints the given messages because it does not take the packageLogLvl into account.
func Log(color colors.Color, args ...interface{}) {
	log(color, args...)
}

// Debug checks whether the packageLogLvl is sufficiently high and calls the logDebug() method from the package if it is.
func Debug(args ...interface{}) {
	if packageLogLvl.lvl >= DebugLvl {
		logDebug(args...)
	}
}

// Info checks whether the packageLogLvl is sufficiently high and calls the logInfo() method from the package if it is.
func Info(args ...interface{}) {
	if packageLogLvl.lvl >= InfoLvl {
		logInfo(args...)
	}
}

// Warn checks whether the packageLogLvl is sufficiently high and calls the logWarn() method from the package if it is.
func Warn(args ...interface{}) {
	if packageLogLvl.lvl >= WarnLvl {
		logWarn(args...)
	}
}

// Error checks whether the packageLogLvl is sufficiently high and calls the logError() method from the package if it is.
func Error(args ...interface{}) {
	if packageLogLvl.lvl >= ErrorLvl {
		logError(args...)
	}
}

// FatalError calls the logFatalError() package method, see its method documentation for more logInfo.
func FatalError(args ...interface{}) {
	logFatalError(args...)
}

// GetLogLvlName returns the package Logger Log Level Name.
func GetLogLvlName() string {
	return packageLogLvl.LvlName()
}

// GetLogLvlIntValue returns the package Logger Log Level int8 value.
func GetLogLvlIntValue() int8 {
	return packageLogLvl.lvl
}

// SetLogLvl updates the package Logger log level property if the given logLvlName is valid,
// otherwise it sets the log level to DebugLvlName.
func SetLogLvl(logLvlName string) {
	packageLogLvl.lvl = retrieveLogLvlIntFromName(logLvlName)
}

// SetLogLvlEnvVariable updates the package Logger log level property attempting to
// retrieve it from the given envVariableName's value.
// If the env variable is not found sets the package log level to DebugLvlName by default.
func SetLogLvlEnvVariable(envVariableName string) {
	packageLogLvl.envVariable = envVariableName
	packageLogLvl.lvl = retrieveLogLvlFromEnv(packageLogLvl.envVariable)
}
