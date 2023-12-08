package logs

type Logger struct {
	logLvl LogLevel
}

// Debug checks whether the instance logLvl is sufficiently high and calls the logDebug() method accordingly.
func (l *Logger) Debug(args ...interface{}) {
	if l.logLvl.lvl >= DebugLvl {
		logDebug(args...)
	}
}

// Info checks whether the instance logLvl is sufficiently high and calls the logInfo() method accordingly.
func (l *Logger) Info(args ...interface{}) {
	if l.logLvl.lvl >= InfoLvl {
		logInfo(args...)
	}
}

// Warn checks whether the instance logLvl is sufficiently high and calls the logWarn() method accordingly.
func (l *Logger) Warn(args ...interface{}) {
	if l.logLvl.lvl >= WarnLvl {
		logWarn(args...)
	}
}

// Error checks whether the instance logLvl is sufficiently high and calls the logError() method accordingly.
func (l *Logger) Error(args ...interface{}) {
	if l.logLvl.lvl >= ErrorLvl {
		logError(args...)
	}
}

// FatalError calls the logFatalError() package method, see its method documentation for more logInfo.
func (l *Logger) FatalError(args ...interface{}) {
	logFatalError(args...)
}

// GetLogLvlName returns the Logger current set Log Level Name.
func (l *Logger) GetLogLvlName() string {
	return logLvlIntToName[l.logLvl.lvl]
}

// GetLogLvlIntValue returns the Logger current set Log Level int8 value.
func (l *Logger) GetLogLvlIntValue() int8 {
	return l.logLvl.lvl
}

// SetLogLvl updates the Logger instance logLvl.lvl property if the given logLvlName is valid,
// otherwise sets the logLvl.lvl to DebugLvlName.
func (l *Logger) SetLogLvl(logLvlName string) *Logger {
	l.logLvl.lvl = retrieveLogLvlIntFromName(logLvlName)

	return l
}

// SetLogLvlEnvVariable updates the Logger instance logLvl.lvl property  attempting to
// retrieve the log level value of the given envVariableName.
// If the env variable is not found sets DebugLvlName.
func (l *Logger) SetLogLvlEnvVariable(envVariableName string) *Logger {
	l.logLvl.envVariable = envVariableName
	l.logLvl.lvl = retrieveLogLvlFromEnv(l.logLvl.envVariable)

	return l
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
