package services

import (
	"bytes"
	"fmt"
	"strconv"
)

// YamlLogEntry represents a structured log entry that can be marshaled to YAML format.
// All fields except Message are optional and will be omitted if empty.
type YamlLogEntry struct {
	Level    string `yaml:"level,omitempty"`
	Date     string `yaml:"date,omitempty"`
	Time     string `yaml:"time,omitempty"`
	DateTime string `yaml:"datetime,omitempty"`
	Message  string `yaml:"msg"`
	Extras   []any  `yaml:"extras,omitempty"`
}

// YamlMarshaler provides custom YAML marshaling functionality optimized for log entries.
type YamlMarshaler struct {
	specialCharsSet map[rune]any
}

// Marshal converts a YamlLogEntry into a YAML-formatted byte slice.
// It uses a buffer-based approach to minimize allocations during marshaling.
func (y *YamlMarshaler) Marshal(logEntry YamlLogEntry) []byte {
	var res bytes.Buffer
	extrasLen := len(logEntry.Extras)
	res.Grow(yamlCharOverhead + (averageExtraLen * extrasLen))

	y.writeLogEntryProperties(&res, logEntry.Level, logEntry.Date, logEntry.Time, logEntry.DateTime)

	res.WriteString("msg: ")
	res.WriteString(logEntry.Message)
	res.WriteByte('\n')

	if extrasLen > 0 {
		res.WriteString("extras:\n")

		for i := 0; i < extrasLen; i += 2 {
			res.WriteString("  ")

			if i < extrasLen {
				y.writeStr(&res, logEntry.Extras[i], true)
				res.WriteString(": ")

				if i+1 < extrasLen {
					y.writeStr(&res, logEntry.Extras[i+1], false)
					res.WriteByte('\n')
				}
			}
		}

		if extrasLen%2 != 0 {
			res.WriteString("null")
			res.WriteByte('\n')
		}
	}

	return res.Bytes()
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
func (y *YamlMarshaler) writeLogEntryProperties(res *bytes.Buffer, level string, date string, time string, dateTime string) {
	if level != "" {
		res.WriteString("level: ")
		res.WriteString(level)
		res.WriteByte('\n')
	}

	if date != "" {
		res.WriteString("date: ")
		res.WriteString(date)
		res.WriteByte('\n')
	}

	if time != "" {
		res.WriteString("time: ")
		res.WriteString(time)
		res.WriteByte('\n')
	}

	if dateTime != "" {
		res.WriteString("datetime: ")
		res.WriteString(dateTime)
		res.WriteByte('\n')
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
