package services

import (
	"bytes"
	"fmt"
	"strconv"

	s "github.com/pho3b/tiny-logger/shared"
)

// YamlLogEntry represents a structured log entry that can be marshaled to YAML format.
// All fields except Message are optional and will be omitted if empty.
type YamlLogEntry struct {
	Level          string           `yaml:"level,omitempty"`
	Date           string           `yaml:"date,omitempty"`
	Time           string           `yaml:"time,omitempty"`
	DateTime       string           `yaml:"datetime,omitempty"`
	DateTimeFormat s.DateTimeFormat `yaml:"dateTimeFormat,omitempty"`
	Message        string           `yaml:"msg"`
	Extras         []any            `yaml:"extras,omitempty"`
}

// YamlMarshaler provides custom YAML marshaling functionality optimized for log entries.
type YamlMarshaler struct {
	specialCharsSet map[rune]any
}

// MarshalInto converts a YamlLogEntry into a YAML-formatted byte slice and adds it to the given buffer
// to minimize allocations during marshaling.
func (y *YamlMarshaler) MarshalInto(buf *bytes.Buffer, logEntry *YamlLogEntry) {
	extrasLen := len(logEntry.Extras)
	buf.Grow(yamlCharOverhead + (averageExtraLen * extrasLen))

	y.writeLogEntryProperties(buf, logEntry.Level, logEntry.Date, logEntry.Time, logEntry.DateTime, logEntry.DateTimeFormat)

	buf.WriteString("msg: ")
	buf.WriteString(logEntry.Message)
	buf.WriteByte('\n')

	if extrasLen > 0 {
		buf.WriteString("extras:\n")

		for i := 0; i < extrasLen; i += 2 {
			buf.WriteString("  ")

			if i < extrasLen {
				y.writeStr(buf, logEntry.Extras[i], true)
				buf.WriteString(": ")

				if i+1 < extrasLen {
					y.writeStr(buf, logEntry.Extras[i+1], false)
					buf.WriteByte('\n')
				}
			}
		}

		if extrasLen%2 != 0 {
			buf.WriteString("null")
			buf.WriteByte('\n')
		}
	}
}

// writeStr writes a string value to the buffer with appropriate YAML formatting.
func (y *YamlMarshaler) writeStr(buf *bytes.Buffer, v any, isKey bool) {
	switch val := v.(type) {
	case string:
		if y.containsSpecialChars(val) {
			buf.WriteByte('"')
			buf.WriteString(val)
			buf.WriteByte('"')
		} else {
			buf.WriteString(val)
		}
	case rune:
		if isKey {
			buf.WriteRune(val)
		} else {
			buf.WriteByte('"')
			buf.WriteRune(val)
			buf.WriteByte('"')
		}
	case int:
		buf.WriteString(strconv.Itoa(val))
	case int64:
		buf.WriteString(strconv.FormatInt(val, 10))
	case float64:
		buf.WriteString(strconv.FormatFloat(val, 'f', -1, 64))
	case bool:
		buf.WriteString(strconv.FormatBool(val))
	default:
		// Check if the string representation needs quotes
		str := fmt.Sprint(val)

		if y.containsSpecialChars(str) {
			buf.WriteByte('"')
			buf.WriteString(str)
			buf.WriteByte('"')
		} else {
			buf.WriteString(str)
		}
	}
}

// writeLogEntryProperties writes the standard log entry properties to the buffer.
// Only non-empty properties are written.
func (y *YamlMarshaler) writeLogEntryProperties(
	buf *bytes.Buffer,
	level string,
	date string,
	time string,
	dateTime string,
	dateTimeFormat s.DateTimeFormat,
) {
	if level != "" {
		buf.WriteString("level: ")
		buf.WriteString(level)
		buf.WriteByte('\n')
	}

	if dateTime != "" || (date != "" && time != "") {
		if dateTimeFormat == s.UnixTimestamp {
			buf.WriteString("ts: ")
		} else {
			buf.WriteString("datetime: ")
		}

		buf.WriteString(dateTime)
		buf.WriteByte('\n')

		return
	}

	if date != "" {
		buf.WriteString("date: ")
		buf.WriteString(date)
		buf.WriteByte('\n')
	}

	if time != "" {
		buf.WriteString("time: ")
		buf.WriteString(time)
		buf.WriteByte('\n')
	}
}

// containsSpecialChars checks if a string contains characters that require quoting in YAML
func (y *YamlMarshaler) containsSpecialChars(s string) bool {
	for _, c := range s {
		if _, ok := y.specialCharsSet[c]; ok {
			return true
		}
	}

	return false
}

func NewYamlMarshaler() YamlMarshaler {
	return YamlMarshaler{
		specialCharsSet: map[rune]any{':': nil, '{': nil, '}': nil, '[': nil, ']': nil, ',': nil, '&': nil,
			'*': nil, '#': nil, '?': nil, '|': nil, '-': nil, '<': nil, '>': nil, '=': nil, '!': nil, '%': nil,
			'@': nil, '`': nil, ' ': nil,
		},
	}
}
