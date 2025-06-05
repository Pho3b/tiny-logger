package logs

import (
	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/encoders"
	"github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/pho3b/tiny-logger/shared"
)

type Logger struct {
	dateEnabled   bool
	timeEnabled   bool
	colorsEnabled bool
	showLogLevel  bool
	encoder       shared.EncoderInterface
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
func (l *Logger) SetLogLvl(logLvlName log_level.LogLvlName) *Logger {
	l.logLvl.Lvl = log_level.RetrieveLogLvlIntFromName(logLvlName)

	return l
}

// SetLogLvlEnvVariable sets the log level based on an environment variable. If the variable is not found,
// defaults to DebugLvlName.
// NOTE: The environment variable value must be a valid log_level.LogLvlName string.
func (l *Logger) SetLogLvlEnvVariable(envVariableName string) *Logger {
	l.logLvl.EnvVariable = envVariableName
	l.logLvl.Lvl = log_level.RetrieveLogLvlFromEnv(l.logLvl.EnvVariable)

	return l
}

func (l *Logger) Color(color colors.Color, args ...interface{}) {
	l.encoder.LogDebug(l, args...)
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
func (l *Logger) EnableColors(enable bool) *Logger {
	l.colorsEnabled = enable

	return l
}

// GetColorsEnabled returns true if color output is enabled, false otherwise.
func (l *Logger) GetColorsEnabled() bool {
	return l.colorsEnabled
}

// ShowLogLevel enables/disables the log level visibility of the logger.
func (l *Logger) ShowLogLevel(enable bool) *Logger {
	l.showLogLevel = enable

	return l
}

// GetShowLogLevel returns the showLogLevel value of the logger.
func (l *Logger) GetShowLogLevel() bool {
	return l.showLogLevel
}

// AddDateTime enables or disables both date and time in log output.
func (l *Logger) AddDateTime(addDateTime bool) *Logger {
	l.dateEnabled = addDateTime
	l.timeEnabled = addDateTime

	return l
}

// AddDate enables or disables date in log output based on the provided parameter.
func (l *Logger) AddDate(addDate bool) *Logger {
	l.dateEnabled = addDate

	return l
}

// AddTime enables or disables time in log output based on the provided parameter.
func (l *Logger) AddTime(addTime bool) *Logger {
	l.timeEnabled = addTime

	return l
}

// GetDateTimeEnabled returns the current date and time settings of the logger.
func (l *Logger) GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool) {
	return l.dateEnabled, l.timeEnabled
}

// SetEncoder sets the Encoder that will be used to print logs.
func (l *Logger) SetEncoder(encoderType shared.EncoderType) *Logger {
	switch encoderType {
	case shared.DefaultEncoderType:
		l.encoder = encoders.NewDefaultEncoder()
	case shared.JsonEncoderType:
		l.encoder = encoders.NewJSONEncoder()
	case shared.YamlEncoderType:
		l.encoder = encoders.NewYAMLEncoder()
	}

	return l
}

// GetEncoderType returns the currently set Encoder type.
func (l *Logger) GetEncoderType() shared.EncoderType {
	return l.encoder.GetType()
}

// NewLogger creates and returns a new Logger instance with default settings.
func NewLogger() *Logger {
	logger := &Logger{
		dateEnabled:   false,
		timeEnabled:   false,
		colorsEnabled: false,
		showLogLevel:  true,
		encoder:       encoders.NewDefaultEncoder(),
	}
	logger.SetLogLvlEnvVariable(log_level.DefaultEnvLogLvlVar)

	return logger
}
