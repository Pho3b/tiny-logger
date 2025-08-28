package encoders

import (
	"bytes"
	"os"
	"sync"

	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
)

type YAMLEncoder struct {
	BaseEncoder
	DateTimePrinter services.DateTimePrinter
	yamlMarshaler   services.YamlMarshaler
}

// LogDebug formats and prints a debug-level log message in YAML format.
func (y *YAMLEncoder) LogDebug(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		y.log(logger, ll.DebugLvlName, s.StdOutput, args...)
	}
}

// LogInfo formats and prints an info-level log message in YAML format.
func (y *YAMLEncoder) LogInfo(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		y.log(logger, ll.InfoLvlName, s.StdOutput, args...)
	}
}

// LogWarn formats and prints a warning-level log message in YAML format.
func (y *YAMLEncoder) LogWarn(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 {
		y.log(logger, ll.WarnLvlName, s.StdOutput, args...)
	}
}

// LogError formats and prints an error-level log message in YAML format.
func (y *YAMLEncoder) LogError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !y.areAllNil(args...) {
		y.log(logger, ll.ErrorLvlName, s.StdErrOutput, args...)
	}
}

// LogFatalError formats and prints a fatal error-level log message in YAML format and exits the program.
func (y *YAMLEncoder) LogFatalError(logger s.LoggerConfigsInterface, args ...any) {
	if len(args) > 0 && !y.areAllNil(args...) {
		y.log(logger, ll.FatalErrorLvlName, s.StdErrOutput, args...)
		os.Exit(1)
	}
}

// Color formats and prints a colored log message using the specified color.
//
// Parameters:
//   - color: the color to apply to the log message.
//   - args: variadic msg arguments.
func (y *YAMLEncoder) Color(lConfig s.LoggerConfigsInterface, color c.Color, args ...any) {
	if len(args) > 0 {
		msgBuffer := y.getBuffer()
		dEnabled, tEnabled := lConfig.GetDateTimeEnabled()
		msgBuffer.WriteString(color.String())

		y.composeMsgInto(
			msgBuffer,
			y.yamlMarshaler,
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			false,
			y.castAndConcatenate(args[0]),
			args[1:]...,
		)

		msgBuffer.WriteString(c.Reset.String())
		y.printLog(s.StdOutput, msgBuffer, true)
		y.putBuffer(msgBuffer)
	}
}

// log formats and prints a log message to the given output type.
// Internally used by all the encoder Log methods.
func (y *YAMLEncoder) log(
	logger s.LoggerConfigsInterface,
	logLvlName ll.LogLvlName,
	outType s.OutputType,
	args ...any,
) {
	dEnabled, tEnabled := logger.GetDateTimeEnabled()
	msgBuffer := y.getBuffer()

	y.composeMsgInto(
		msgBuffer,
		y.yamlMarshaler,
		logLvlName,
		dEnabled,
		tEnabled,
		logger.GetShowLogLevel(),
		y.castAndConcatenate(args[0]),
		args[1:]...,
	)

	y.printLog(outType, msgBuffer, true)
	y.putBuffer(msgBuffer)
}

// composeMsgInto formats and writes the given 'msg' into the given buffer.
func (y *YAMLEncoder) composeMsgInto(
	buf *bytes.Buffer,
	yamlMarshaler services.YamlMarshaler,
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	showLogLevel bool,
	msg string,
	extras ...any,
) {
	buf.Grow((averageWordLen * len(extras)) + len(msg) + 60)
	date, time, dateTime := y.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)

	if !showLogLevel {
		logLevel = ""
	}

	yamlMarshaler.MarshalInto(
		buf,
		services.YamlLogEntry{
			Level:    logLevel.String(),
			Date:     date,
			Time:     time,
			DateTime: dateTime,
			Message:  msg,
			Extras:   extras,
		},
	)
}

// NewYAMLEncoder initializes and returns a new YAMLEncoder instance.
func NewYAMLEncoder() *YAMLEncoder {
	encoder := &YAMLEncoder{DateTimePrinter: services.NewDateTimePrinter(), yamlMarshaler: services.NewYamlMarshaler()}
	encoder.encoderType = s.YamlEncoderType
	encoder.bufferSyncPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	return encoder
}
