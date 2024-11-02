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

// Debug checks whether the instance logLvl is sufficiently high and calls the logDebug() method accordingly.
func (l *Logger) Debug(args ...interface{}) {
	if l.logLvl.Lvl >= log_level.DebugLvl {
		l.encoder.LogDebug(l, args...)
	}
}

// Info checks whether the instance logLvl is sufficiently high and calls the logInfo() method accordingly.
func (l *Logger) Info(args ...interface{}) {
	if l.logLvl.Lvl >= log_level.InfoLvl {
		l.encoder.LogInfo(l, args...)
	}
}

// Warn checks whether the instance logLvl is sufficiently high and calls the logWarn() method accordingly.
func (l *Logger) Warn(args ...interface{}) {
	if l.logLvl.Lvl >= log_level.WarnLvl {
		l.encoder.LogWarn(l, args...)
	}
}

// Error checks whether the instance logLvl is sufficiently high and calls the logError() method accordingly.
func (l *Logger) Error(args ...interface{}) {
	if l.logLvl.Lvl >= log_level.ErrorLvl {
		l.encoder.LogError(l, args...)
	}
}

// FatalError calls the logFatalError() package method, see its method documentation for more logInfo.
func (l *Logger) FatalError(args ...interface{}) {
	l.encoder.LogFatalError(l, args...)
}

func (l *Logger) GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool) {
	return l.dateEnabled, l.timeEnabled
}

func (l *Logger) GetColorsEnabled() bool {
	return l.colorsEnabled
}

// GetLogLvlName returns the Logger current set Log Level Name.
func (l *Logger) GetLogLvlName() log_level.LogLvlName {
	return log_level.LogLvlIntToName[l.logLvl.Lvl]
}

// GetLogLvlIntValue returns the Logger current set Log Level int8 value.
func (l *Logger) GetLogLvlIntValue() int8 {
	return l.logLvl.Lvl
}

// SetLogLvl updates the Logger instance logLvl.Lvl property if the given logLvlName is valid,
// otherwise sets the logLvl.Lvl to DebugLvlName.
func (l *Logger) SetLogLvl(logLvlName log_level.LogLvlName) interfaces.LoggerInterface {
	l.logLvl.Lvl = log_level.RetrieveLogLvlIntFromName(logLvlName)

	return l
}

func (l *Logger) EnableColors(enable bool) interfaces.LoggerInterface {
	l.colorsEnabled = enable

	return l
}

func (l *Logger) AddDateTime(addDateTime bool) interfaces.LoggerInterface {
	l.dateEnabled = addDateTime
	l.timeEnabled = addDateTime

	return l
}

func (l *Logger) AddDate(addDate bool) interfaces.LoggerInterface {
	l.dateEnabled = addDate

	return l
}

func (l *Logger) AddTime(addTime bool) interfaces.LoggerInterface {
	l.timeEnabled = addTime

	return l
}

// SetLogLvlEnvVariable updates the Logger instance logLvl.Lvl property  attempting to
// retrieve the log level value of the given envVariableName.
// If the env variable is not found sets DebugLvlName by default.
func (l *Logger) SetLogLvlEnvVariable(envVariableName string) interfaces.LoggerInterface {
	l.logLvl.EnvVariable = envVariableName
	l.logLvl.Lvl = log_level.RetrieveLogLvlFromEnv(l.logLvl.EnvVariable)

	return l
}

// NewLogger returns a new logger with the logLvl set to 'DebugLvl' by default.
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
