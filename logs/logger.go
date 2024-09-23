package logs

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/internal/shared"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/configs"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type Logger struct {
	conf *configs.TLConfigs
}

// Log calls the underlying log() method from the package.
// It always prints the given messages because it does not take the packageLogLvl into account.
func (l *Logger) Log(color colors.Color, args ...interface{}) {
	shared.Log(color, args...)
}

// Debug checks whether the instance logLvl is sufficiently high and calls the logDebug() method accordingly.
func (l *Logger) Debug(args ...interface{}) {
	if l.conf.LogLvl.Lvl >= log_level.DebugLvl {
		shared.LogDebug(l.conf, args...)
	}
}

// Info checks whether the instance logLvl is sufficiently high and calls the logInfo() method accordingly.
func (l *Logger) Info(args ...interface{}) {
	if l.conf.LogLvl.Lvl >= log_level.InfoLvl {
		shared.LogInfo(args...)
	}
}

// Warn checks whether the instance logLvl is sufficiently high and calls the logWarn() method accordingly.
func (l *Logger) Warn(args ...interface{}) {
	if l.conf.LogLvl.Lvl >= log_level.WarnLvl {
		shared.LogWarn(args...)
	}
}

// Error checks whether the instance logLvl is sufficiently high and calls the logError() method accordingly.
func (l *Logger) Error(args ...interface{}) {
	if l.conf.LogLvl.Lvl >= log_level.ErrorLvl {
		shared.LogError(args...)
	}
}

// FatalError calls the logFatalError() package method, see its method documentation for more logInfo.
func (l *Logger) FatalError(args ...interface{}) {
	shared.LogFatalError(args...)
}

// GetLogLvlName returns the Logger current set Log Level Name.
func (l *Logger) GetLogLvlName() log_level.LogLvlName {
	return log_level.LogLvlIntToName[l.conf.LogLvl.Lvl]
}

// GetLogLvlIntValue returns the Logger current set Log Level int8 value.
func (l *Logger) GetLogLvlIntValue() int8 {
	return l.conf.LogLvl.Lvl
}

// SetLogLvl updates the Logger instance logLvl.Lvl property if the given logLvlName is valid,
// otherwise sets the logLvl.Lvl to DebugLvlName.
func (l *Logger) SetLogLvl(logLvlName log_level.LogLvlName) *Logger {
	l.conf.LogLvl.Lvl = log_level.RetrieveLogLvlIntFromName(logLvlName)

	return l
}

// SetLogLvlEnvVariable updates the Logger instance logLvl.Lvl property  attempting to
// retrieve the log level value of the given envVariableName.
// If the env variable is not found sets DebugLvlName.
func (l *Logger) SetLogLvlEnvVariable(envVariableName string) *Logger {
	l.conf.LogLvl.EnvVariable = envVariableName
	l.conf.LogLvl.Lvl = log_level.RetrieveLogLvlFromEnv(l.conf.LogLvl.EnvVariable)

	return l
}

func (l *Logger) SetConfigs(tlConfigs *configs.TLConfigs) *Logger {
	l.conf = tlConfigs

	return l
}

// NewLogger returns a new logger with the logLvl set to 'DebugLvl' by default.
func NewLogger() *Logger {
	return &Logger{
		conf: configs.NewDefaultTLConfigs(),
	}
}
