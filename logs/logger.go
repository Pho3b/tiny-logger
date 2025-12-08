package logs

import (
	"os"
	"time"

	"github.com/pho3b/tiny-logger/internal/services"
	"github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/encoders"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
)

type Logger struct {
	dateEnabled    bool
	timeEnabled    bool
	colorsEnabled  bool
	showLogLevel   bool
	encoder        s.EncoderInterface
	logLvl         ll.LogLevel
	outFile        *os.File
	dateTimeFormat s.DateTimeFormat
}

// Debug logs a debug-level message if the logger's log level allows it.
func (l *Logger) Debug(args ...any) {
	if l.logLvl.Lvl >= ll.DebugLvl && len(args) > 0 {
		l.encoder.Log(l, ll.DebugLvlName, l.checkOutFile(s.StdOutput), args...)
	}
}

// Info logs an informational-level message if the logger's log level allows it.
func (l *Logger) Info(args ...any) {
	if l.logLvl.Lvl >= ll.InfoLvl && len(args) > 0 {
		l.encoder.Log(l, ll.InfoLvlName, l.checkOutFile(s.StdOutput), args...)
	}
}

// Warn logs a warning-level message if the logger's log level allows it.
func (l *Logger) Warn(args ...any) {
	if l.logLvl.Lvl >= ll.WarnLvl && len(args) > 0 {
		l.encoder.Log(l, ll.WarnLvlName, l.checkOutFile(s.StdOutput), args...)
	}
}

// Error logs an error-level message if the logger's log level allows it.
func (l *Logger) Error(args ...any) {
	if l.logLvl.Lvl >= ll.ErrorLvl && len(args) > 0 && !l.areAllNil(args...) {
		l.encoder.Log(l, ll.ErrorLvlName, l.checkOutFile(s.StdErrOutput), args...)
	}
}

// FatalError logs a fatal error message and terminates the application only if any given args is not NIl,
// otherwise the method does nothing.
func (l *Logger) FatalError(args ...any) {
	if l.logLvl.Lvl >= ll.ErrorLvl && len(args) > 0 && !l.areAllNil(args...) {
		l.encoder.Log(l, ll.FatalErrorLvlName, l.checkOutFile(s.StdErrOutput), args...)
		os.Exit(1)
	}
}

// Color formats and prints a colored log message using the specified color.
func (l *Logger) Color(color colors.Color, args ...any) {
	l.encoder.Color(l, color, args...)
}

// GetLogLvlName returns the current log level name as a string.
func (l *Logger) GetLogLvlName() ll.LogLvlName {
	return ll.LogLvlIntToName[l.logLvl.Lvl]
}

// GetLogLvlIntValue returns the current log level as an int8 value.
func (l *Logger) GetLogLvlIntValue() int8 {
	return l.logLvl.Lvl
}

// SetLogLvl sets the log level of the logger based on a provided log level name.
// If the provided name is invalid, it defaults to DebugLvlName.
func (l *Logger) SetLogLvl(logLvlName ll.LogLvlName) *Logger {
	l.logLvl.Lvl = ll.RetrieveLogLvlIntFromName(logLvlName)

	return l
}

// SetLogLvlEnvVariable sets the log level based on an environment variable. If the variable is not found,
// defaults to DebugLvlName.
//
// NOTE: The environment variable value must be a valid ll.LogLvlName string.
func (l *Logger) SetLogLvlEnvVariable(envVariableName string) *Logger {
	l.logLvl.EnvVariable = envVariableName
	l.logLvl.Lvl = ll.RetrieveLogLvlFromEnv(l.logLvl.EnvVariable)

	return l
}

// GetColorsEnabled returns true if color output is enabled, false otherwise.
func (l *Logger) GetColorsEnabled() bool {
	return l.colorsEnabled
}

// EnableColors enables or disables color output in the logger based on the given parameter.
// Colors apply only on the header elements [Data, Time, Log Level]
func (l *Logger) EnableColors(enable bool) *Logger {
	l.colorsEnabled = enable

	return l
}

// GetShowLogLevel returns the showLogLevel value of the logger.
func (l *Logger) GetShowLogLevel() bool {
	return l.showLogLevel
}

// ShowLogLevel enables/disables the log level visibility of the logger.
func (l *Logger) ShowLogLevel(enable bool) *Logger {
	l.showLogLevel = enable

	return l
}

// GetDateTimeEnabled returns the current date and time settings of the logger.
func (l *Logger) GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool) {
	return l.dateEnabled, l.timeEnabled
}

// AddDateTime enables or disables both date and time in log output.
func (l *Logger) AddDateTime(addDateTime bool) *Logger {
	l.dateEnabled = addDateTime
	l.timeEnabled = addDateTime

	return l
}

// AddDate enables or disables the date in log output based on the provided parameter.
func (l *Logger) AddDate(addDate bool) *Logger {
	l.dateEnabled = addDate

	return l
}

// AddTime enables or disables time in log output based on the provided parameter.
func (l *Logger) AddTime(addTime bool) *Logger {
	l.timeEnabled = addTime

	return l
}

// GetEncoderType returns the currently set Encoder type.
func (l *Logger) GetEncoderType() s.EncoderType {
	return l.encoder.GetType()
}

// SetEncoder sets the Encoder that will be used to print logs.
func (l *Logger) SetEncoder(encoderType s.EncoderType) *Logger {
	switch encoderType {
	case s.DefaultEncoderType:
		l.encoder = encoders.NewDefaultEncoder()
	case s.JsonEncoderType:
		l.encoder = encoders.NewJSONEncoder()
	case s.YamlEncoderType:
		l.encoder = encoders.NewYAMLEncoder()
	}

	l.encoder.SetDateTimeFormat(l.dateTimeFormat)
	return l
}

// GetLogFile returns the current log file. If no file is set, it returns nil.
func (l *Logger) GetLogFile() *os.File {
	return l.outFile
}

// SetLogFile sets the given os.File as the current Logger output log file.
// If the given file is nil, a warning is logged and the method does nothing.
func (l *Logger) SetLogFile(file *os.File) *Logger {
	if file == nil {
		l.Warn("the given log file is nil, skipping logs redirection")
		return l
	}

	l.outFile = file
	return l
}

// CloseLogFile closes the current log file if it exists. If no file is set, a warning is logged
// and the method does nothing.
func (l *Logger) CloseLogFile() error {
	if l.outFile == nil {
		l.Warn("no log file opened, skipping close")
		return nil
	}

	if err := l.outFile.Close(); err != nil {
		return err
	}

	l.outFile = nil
	return nil
}

// GetDateTimeFormat returns the current DateTimeFormat of the logger.
func (l *Logger) GetDateTimeFormat() s.DateTimeFormat {
	return l.dateTimeFormat
}

// SetDateTimeFormat sets the DateTimeFormat of the logger.
func (l *Logger) SetDateTimeFormat(format s.DateTimeFormat) *Logger {
	l.dateTimeFormat = format
	l.encoder.SetDateTimeFormat(l.dateTimeFormat)

	return l
}

func (l *Logger) EnableBufferedLogs(_ time.Duration) *Logger {
	buf := services.NewLogsBuffer(l)

	go func() {
		time.Sleep(25 * time.Second)
		buf.StopLogs()
	}()

	return l
}

func (l *Logger) GetBufferFlushInterval() time.Duration {
	return 3000 * time.Millisecond
}

// areAllNil returns true if all the given args are 'nil', false otherwise.
func (l *Logger) areAllNil(args ...any) bool {
	for _, arg := range args {
		if arg != nil {
			return false
		}
	}

	return true
}

// checkOutFile returns FileOutput if a log file is set, otherwise returns the provided outType.
func (l *Logger) checkOutFile(outType s.OutputType) s.OutputType {
	if l.outFile != nil {
		return s.FileOutput
	}

	return outType
}

// NewLogger creates and returns a new Logger instance with default settings.
func NewLogger() *Logger {
	logger := &Logger{showLogLevel: true, dateTimeFormat: s.IT}
	logger.SetLogLvlEnvVariable(ll.DefaultEnvLogLvlVar)
	logger.SetEncoder(s.DefaultEncoderType)

	return logger
}
