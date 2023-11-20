package logs

// packageLogLvl represents the package logger log level.
var packageLogLvl = LogLevel{
	lvl:         retrieveLogLvlIntFromName(DebugLvlName),
	envVariable: "",
}

// Debug checks whether the packageLogLvl is sufficiently high and calls the debug() method from the package if it is.
func Debug(args ...interface{}) {
	if packageLogLvl.lvl >= DebugLvl {
		debug(args...)
	}
}

// Info checks whether the packageLogLvl is sufficiently high and calls the info() method from the package if it is.
func Info(args ...interface{}) {
	if packageLogLvl.lvl >= InfoLvl {
		info(args...)
	}
}

// Warn checks whether the packageLogLvl is sufficiently high and calls the warn() method from the package if it is.
func Warn(args ...interface{}) {
	if packageLogLvl.lvl >= WarnLvl {
		warn(args...)
	}
}

// Error checks whether the packageLogLvl is sufficiently high and calls the error() method from the package if it is.
func Error(args ...interface{}) {
	if packageLogLvl.lvl >= ErrorLvl {
		error(args...)
	}
}

// FatalError calls the fatalError() package method, see its method documentation for more info.
func FatalError(args ...interface{}) {
	fatalError(args...)
}

// LogLvlName returns the package Logger Log Level Name.
func LogLvlName() string {
	return packageLogLvl.LvlName()
}

// LogLvlIntValue returns the package Logger Log Level int8 value.
func LogLvlIntValue() int8 {
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
