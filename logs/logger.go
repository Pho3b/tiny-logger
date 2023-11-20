package logs

type Logger struct {
	logLvl LogLevel
}

// Debug checks whether the instance logLvl is sufficiently high and calls the debug() method accordingly.
func (l *Logger) Debug(args ...interface{}) {
	if l.logLvl.lvl >= DebugLvl {
		debug(args...)
	}
}

// Info checks whether the instance logLvl is sufficiently high and calls the info() method accordingly.
func (l *Logger) Info(args ...interface{}) {
	if l.logLvl.lvl >= InfoLvl {
		info(args...)
	}
}

// Warn checks whether the instance logLvl is sufficiently high and calls the warn() method accordingly.
func (l *Logger) Warn(args ...interface{}) {
	if l.logLvl.lvl >= WarnLvl {
		warn(args...)
	}
}

// Error checks whether the instance logLvl is sufficiently high and calls the error() method accordingly.
func (l *Logger) Error(args ...interface{}) {
	if l.logLvl.lvl >= ErrorLvl {
		error(args...)
	}
}

// FatalError calls the fatalError() package method, see its method documentation for more info.
func (l *Logger) FatalError(args ...interface{}) {
	fatalError(args...)
}

// LogLvlName returns the Logger current set Log Level Name.
func (l *Logger) LogLvlName() string {
	return logLvlIntToName[l.logLvl.lvl]
}

// LogLvlIntValue returns the Logger current set Log Level int8 value.
func (l *Logger) LogLvlIntValue() int8 {
	return l.logLvl.lvl
}

// SetLogLvl updates the Logger instance logLvl.lvl property if the given logLvlName is valid,
// otherwise sets the logLvl.lvl to DebugLvlName.
func (l *Logger) SetLogLvl(logLvlName string) {
	l.logLvl.lvl = retrieveLogLvlIntFromName(logLvlName)
}

// SetLogLvlEnvVariable updates the Logger instance logLvl.lvl property attempting to
// retrieve the log level value of the given envVariableName.
// If the env variable is not found sets DebugLvlName.
func (l *Logger) SetLogLvlEnvVariable(envVariableName string) {
	l.logLvl.envVariable = envVariableName
	l.logLvl.lvl = retrieveLogLvlFromEnv(l.logLvl.envVariable)
}

// NewLogger returns a new logger with the logLvl set to 'DebugLvl' by default.
func NewLogger() *Logger {
	return &Logger{
		logLvl: LogLevel{
			lvl:         retrieveLogLvlIntFromName(DebugLvlName),
			envVariable: "",
		},
	}
}
