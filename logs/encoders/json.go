package encoders

import (
	"encoding/json"
	"fmt"
	"github.com/Pho3b/tiny-logger/internal/services"
	"github.com/Pho3b/tiny-logger/shared"
	"os"
)

type JSONEncoder struct {
	BaseEncoder
	DateTimePrinter *services.DateTimePrinter
}

// jsonLogEntry represents the structure of a JSON log entry.
type jsonLogEntry struct {
	Level   string `json:"level"`
	Date    string `json:"date,omitempty"`
	Time    string `json:"time,omitempty"`
	Message string `json:"message"`
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
	if len(args) > 0 {
		j.printJSONLog("ERROR", logger, shared.StdErrOutput, args...)
	}
}

// LogFatalError formats and prints a fatal error-level log message in JSON format and exits the program.
func (j *JSONEncoder) LogFatalError(logger shared.LoggerConfigsInterface, args ...interface{}) {
	if len(args) > 0 {
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
	dateStr, timeStr := j.DateTimePrinter.PrintDateTime(logger.GetDateTimeEnabled())

	msgBytes, err := json.Marshal(
		jsonLogEntry{
			Level:   level,
			Date:    dateStr,
			Time:    timeStr,
			Message: j.buildMsg(args...),
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

// NewJSONEncoder initializes and returns a new JSONEncoder instance.
func NewJSONEncoder() *JSONEncoder {
	encoder := &JSONEncoder{
		DateTimePrinter: services.NewDateTimePrinter(),
	}
	encoder.encoderType = shared.JsonEncoderType

	return encoder
}
