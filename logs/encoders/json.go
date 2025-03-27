package encoders

import (
	"encoding/json"
	"fmt"
	"github.com/pho3b/tiny-logger/internal/services"
	"github.com/pho3b/tiny-logger/shared"
	"os"
)

type JSONEncoder struct {
	BaseEncoder
	DateTimePrinter *services.DateTimePrinter
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
func (j *JSONEncoder) LogDebug(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		j.printJSONLog("DEBUG", logger, shared.StdOutput, args...)
	}
}

// LogInfo formats and prints an info-level log message in JSON format.
func (j *JSONEncoder) LogInfo(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		j.printJSONLog("INFO", logger, shared.StdOutput, args...)
	}
}

// LogWarn formats and prints a warning-level log message in JSON format.
func (j *JSONEncoder) LogWarn(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
		j.printJSONLog("WARN", logger, shared.StdOutput, args...)
	}
}

// LogError formats and prints an error-level log message in JSON format.
func (j *JSONEncoder) LogError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 && !j.areAllNil(args...) {
		j.printJSONLog("ERROR", logger, shared.StdErrOutput, args...)
	}
}

// LogFatalError formats and prints a fatal error-level log message in JSON format and exits the program.
func (j *JSONEncoder) LogFatalError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 && !j.areAllNil(args...) {
		j.printJSONLog("FATAL", logger, shared.StdErrOutput, args...)
		os.Exit(1)
	}
}

// printJSONLog formats a log message as JSON and prints it to the appropriate output (stdout or stderr).
func (j *JSONEncoder) printJSONLog(
	level string,
	logger shared.LoggerConfigsInterface,
	outType shared.OutputType,
	args ...interface{},
) {
	dateStr, timeStr, dateTimeStr := j.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())

	if !logger.GetShowLogLevel() {
		level = ""
	}

	msgBytes, err := json.Marshal(
		jsonLogEntry{
			Level:    level,
			Date:     dateStr,
			DateTime: dateTimeStr,
			Time:     timeStr,
			Message:  j.buildMsg(args[0]),
			Extras:   buildExtraMessages(args[1:]...),
		},
	)
	msgBytes = append(msgBytes, '\n')
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `{"level":"ERROR", "message":"Failed to marshal JSON log: %s"}`, err.Error())
		return
	}

	switch outType {
	case shared.StdOutput:
		_, _ = os.Stdout.Write(msgBytes)
	case shared.StdErrOutput:
		_, _ = os.Stderr.Write(msgBytes)
	}
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
func buildExtraMessages(keyAndValuePairs ...interface{}) map[string]interface{} {
	keyAndValuePairsLen := len(keyAndValuePairs)
	if keyAndValuePairsLen == 0 {
		return nil
	}

	resMap := make(map[string]interface{}, keyAndValuePairsLen/2)

	for i := 0; i < keyAndValuePairsLen-1; i += 2 {
		key := fmt.Sprint(keyAndValuePairs[i])
		value := keyAndValuePairs[i+1]
		resMap[key] = value
	}

	if keyAndValuePairsLen%2 != 0 {
		lastKey := fmt.Sprint(keyAndValuePairs[keyAndValuePairsLen-1])
		resMap[lastKey] = nil
	}

	return resMap
}

// NewJSONEncoder initializes and returns a new JSONEncoder instance.
func NewJSONEncoder() *JSONEncoder {
	encoder := &JSONEncoder{
		DateTimePrinter: services.NewDateTimePrinter(),
	}
	encoder.encoderType = shared.JsonEncoderType

	return encoder
}
