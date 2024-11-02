package logs

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/interfaces"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/encoders"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type Logger struct {
	dateEnabled   bool
	timeEnabled   bool
	colorsEnabled bool
	encoder       interfaces.Encoder
	logLvl        log_level.LogLevel
}

// Debug logs a debug-level message if the logger's log level allows it.
func (l *Logger) Debug(args ...interface{}) {
	if l.logLvl.Lvl >= log_level.DebugLvl {
		l.encoder.LogDebug(l, args...)
	}
}

// Info logs an informational-level message if the logger's log level allows it.
func (l *Logger) Info(args ...interface{}) {
	if l.logLvl.Lvl >= log_level.InfoLvl {
		l.encoder.LogInfo(l, args...)
	}
}

// Warn logs a warning-level message if the logger's log level allows it.
func (l *Logger) Warn(args ...interface{}) {
	if l.logLvl.Lvl >= log_level.WarnLvl {
		l.encoder.LogWarn(l, args...)
	}
}

// Error logs an error-level message if the logger's log level allows it.
func (l *Logger) Error(args ...interface{}) {
	if l.logLvl.Lvl >= log_level.ErrorLvl {
		l.encoder.LogError(l, args...)
	}
}

// FatalError logs a fatal error message and typically terminates the application.
func (l *Logger) FatalError(args ...interface{}) {
	l.encoder.LogFatalError(l, args...)
}

// SetLogLvl sets the log level of the logger based on a provided log level name.
// If the provided name is invalid, it defaults to DebugLvlName.
func (l *Logger) SetLogLvl(logLvlName log_level.LogLvlName) interfaces.LoggerInterface {
	l.logLvl.Lvl = log_level.RetrieveLogLvlIntFromName(logLvlName)

	return l
}

// SetLogLvlEnvVariable sets the log level based on an environment variable. If the variable is not found,
// defaults to DebugLvlName.
// NOTE: The environment variable value must be a valid log_level.LogLvlName string.
func (l *Logger) SetLogLvlEnvVariable(envVariableName string) interfaces.LoggerInterface {
	l.logLvl.EnvVariable = envVariableName
	l.logLvl.Lvl = log_level.RetrieveLogLvlFromEnv(l.logLvl.EnvVariable)

	return l
}

// GetLogLvlName returns the current log level name as a string.
func (l *Logger) GetLogLvlName() log_level.LogLvlName {
	return log_level.LogLvlIntToName[l.logLvl.Lvl]
}

// GetLogLvlIntValue returns the current log level as an int8 value.
func (l *Logger) GetLogLvlIntValue() int8 {
	return l.logLvl.Lvl
}

// EnableColors enables or disables color output in the logger based on the given parameter.
func (l *Logger) EnableColors(enable bool) interfaces.LoggerInterface {
	l.colorsEnabled = enable

	return l
}

// GetColorsEnabled returns true if color output is enabled, false otherwise.
func (l *Logger) GetColorsEnabled() bool {
	return l.colorsEnabled
}

// AddDateTime enables or disables both date and time in log output.
func (l *Logger) AddDateTime(addDateTime bool) interfaces.LoggerInterface {
	l.dateEnabled = addDateTime
	l.timeEnabled = addDateTime

	return l
}

// AddDate enables or disables date in log output based on the provided parameter.
func (l *Logger) AddDate(addDate bool) interfaces.LoggerInterface {
	l.dateEnabled = addDate

	return l
}

// AddTime enables or disables time in log output based on the provided parameter.
func (l *Logger) AddTime(addTime bool) interfaces.LoggerInterface {
	l.timeEnabled = addTime

	return l
}

// GetDateTimeEnabled returns the current date and time settings of the logger.
func (l *Logger) GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool) {
	return l.dateEnabled, l.timeEnabled
}

// NewLogger creates and returns a new Logger instance with default settings.
func NewLogger() *Logger {
	logger := &Logger{
		dateEnabled:   false,
		timeEnabled:   false,
		colorsEnabled: false,
		encoder:       encoders.NewDefaultEncoder(),
	}
	logger.SetLogLvlEnvVariable(log_level.DefaultEnvLogLvlVar)

	return logger
}
