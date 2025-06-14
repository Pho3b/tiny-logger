package encoders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pho3b/tiny-logger/internal/services"
	c "github.com/pho3b/tiny-logger/logs/colors"
	ll "github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
	"os"
)

type JSONEncoder struct {
	BaseEncoder
	DateTimePrinter services.DateTimePrinter
}

// jsonLogEntry represents the structure of a JSON log entry.
type jsonLogEntry struct {
	Level    string                 `json:"level,omitempty"`
	Date     string                 `json:"date,omitempty"`
	Time     string                 `json:"time,omitempty"`
	DateTime string                 `json:"datetime,omitempty"`
	Message  string                 `json:"msg"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

// LogDebug formats and prints a debug-level log message in JSON format.
func (j *JSONEncoder) LogDebug(logger s.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			ll.DebugLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogInfo formats and prints an info-level log message in JSON format.
func (j *JSONEncoder) LogInfo(logger s.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogWarn formats and prints a warning-level log message in JSON format.
func (j *JSONEncoder) LogWarn(logger s.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			ll.WarnLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdOutput, msgBuffer, true)
	}
}

// LogError formats and prints an error-level log message in JSON format.
func (j *JSONEncoder) LogError(logger s.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 && !j.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			ll.ErrorLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdErrOutput, msgBuffer, true)
	}
}

// LogFatalError formats and prints a fatal error-level log message in JSON format and exits the program.
func (j *JSONEncoder) LogFatalError(logger s.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 && !j.areAllNil(args...) {
		dEnabled, tEnabled := logger.GetDateTimeEnabled()
		msgBuffer := j.composeMsg(
			ll.FatalErrorLvlName,
			dEnabled,
			tEnabled,
			logger.GetShowLogLevel(),
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		j.printLog(s.StdErrOutput, msgBuffer, true)
		os.Exit(1)
	}
}

// Color formats and prints a colored log message using the specified color.
//
// Parameters:
//   - color: the color to apply to the log message.
//   - args: variadic arguments where the first is treated as the message and the rest are appended.
func (j *JSONEncoder) Color(lConfig s.LoggerConfigsInterface, color c.Color, args ...interface{}) {
	if len(args) > 0 {
		var b bytes.Buffer
		b.Grow((len(args) * averageWordLen) + averageWordLen)
		dEnabled, tEnabled := lConfig.GetDateTimeEnabled()

		msgBuffer := j.composeMsg(
			ll.InfoLvlName,
			dEnabled,
			tEnabled,
			false,
			j.castAndConcatenate(args[0]),
			args[1:]...,
		)

		b.WriteString(color.String())
		b.Write(msgBuffer.Bytes())
		b.WriteString(c.Reset.String())

		j.printLog(s.StdOutput, b, true)
	}
}

func (j *JSONEncoder) composeMsg(
	logLevel ll.LogLvlName,
	dateEnabled bool,
	timeEnabled bool,
	showLogLevel bool,
	msg string,
	extras ...interface{},
) bytes.Buffer {
	var b bytes.Buffer
	b.Grow((averageWordLen * len(extras)) + len(msg) + 60)
	dateStr, timeStr, dateTimeStr := j.DateTimePrinter.RetrieveDateTime(dateEnabled, timeEnabled)

	if !showLogLevel {
		logLevel = ""
	}

	msgBytes, err := json.Marshal(
		jsonLogEntry{
			Level:    logLevel.String(),
			Date:     dateStr,
			DateTime: dateTimeStr,
			Time:     timeStr,
			Message:  msg,
			Extras:   j.buildExtraMessages(extras...),
		},
	)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `{"level":"ERROR", "message":"Failed to marshal JSON log: %s"}`, err.Error())
		return bytes.Buffer{}
	}

	b.Write(msgBytes)
	return b
}

// buildExtraMessages constructs a map from a variadic list of key-value pairs.
// It expects an even number of arguments, where even indices (0, 2, 4, ...) are keys
// and odd indices (1, 3, 5, ...) are values. If an odd number of arguments is passed,
// the last key will be assigned a `nil` value.
//
// Example Usage:
//
//	extra := b.buildExtraMessages("user", "alice", "ip", "192.168.1.1")
//	// Result: map[string]interface{}{"user": "alice", "ip": "192.168.1.1"}
func (j *JSONEncoder) buildExtraMessages(keyAndValuePairs ...interface{}) map[string]interface{} {
	keyAndValuePairsLen := len(keyAndValuePairs)
	if keyAndValuePairsLen == 0 {
		return nil
	}

	resMap := make(map[string]interface{}, keyAndValuePairsLen/2)

	for i := 0; i < keyAndValuePairsLen-1; i += 2 {
		key := j.castAndConcatenate(keyAndValuePairs[i])
		value := keyAndValuePairs[i+1]
		resMap[key] = value
	}

	if keyAndValuePairsLen%2 != 0 {
		lastKey := j.castAndConcatenate(keyAndValuePairs[keyAndValuePairsLen-1])
		resMap[lastKey] = nil
	}

	return resMap
}

// NewJSONEncoder initializes and returns a new JSONEncoder instance.
func NewJSONEncoder() *JSONEncoder {
	encoder := &JSONEncoder{
		DateTimePrinter: services.NewDateTimePrinter(),
	}
	encoder.encoderType = s.JsonEncoderType

	return encoder
}
