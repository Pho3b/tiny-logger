package services

import (
	"bytes"
	"fmt"
	"strconv"
)

const yamlCharOverhead = 70

// YamlLogEntry represents a structured log entry that can be marshaled to YAML format.
// All fields except Message are optional and will be omitted if empty.
type YamlLogEntry struct {
	Level   string `yaml:"level,omitempty"`
	Date    string `yaml:"date,omitempty"`
	Time    string `yaml:"time,omitempty"`
	UnixTS  string `yaml:"ts,omitempty"`
	Message string `yaml:"msg"`
	Extras  []any  `yaml:"extras,omitempty"`
}

// YamlMarshaler provides custom YAML marshaling functionality optimized for log entries.
type YamlMarshaler struct {
	specialCharsSet map[rune]any
}

// MarshalInto converts a YamlLogEntry into a YAML-formatted byte slice and adds it to the given buffer
// to minimize allocations during marshaling.
func (y *YamlMarshaler) MarshalInto(buf *bytes.Buffer, logEntry YamlLogEntry) {
	extrasLen := len(logEntry.Extras)
	buf.Grow(yamlCharOverhead + (averageExtraLen * extrasLen))

	y.writeLogEntryProperties(buf, logEntry.Level, logEntry.Date, logEntry.Time, logEntry.UnixTS)

	buf.WriteString("msg: ")
	y.writeScalar(buf, logEntry.Message)
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
			writeYAMLEscaped(buf, val)
			buf.WriteByte('"')
		} else {
			buf.WriteString(val)
		}
	case rune:
		if isKey {
			buf.WriteRune(val)
		} else {
			buf.WriteByte('"')
			writeYAMLEscaped(buf, string(val))
			buf.WriteByte('"')
		}
	case int:
		buf.Write(strconv.AppendInt(buf.AvailableBuffer(), int64(val), 10))
	case int64:
		buf.Write(strconv.AppendInt(buf.AvailableBuffer(), val, 10))
	case float64:
		buf.Write(strconv.AppendFloat(buf.AvailableBuffer(), val, 'f', -1, 64))
	case bool:
		buf.Write(strconv.AppendBool(buf.AvailableBuffer(), val))
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
	unixTS string,
) {
	if level != "" {
		buf.WriteString("level: ")
		buf.WriteString(level)
		buf.WriteByte('\n')
	}

	if unixTS != "" {
		buf.WriteString("ts: ")
		buf.WriteString(unixTS)
		buf.WriteByte('\n')

		return
	}

	if date != "" && time != "" {
		buf.WriteString("datetime: ")
		buf.WriteString(date)
		buf.WriteByte(' ')
		buf.WriteString(time)
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

// writeScalar writes a plain YAML scalar, double-quoting and escaping it only when leaving it
// unquoted would make the YAML invalid (e.g. it contains a newline, a quote or a ': ' sequence).
func (y *YamlMarshaler) writeScalar(buf *bytes.Buffer, s string) {
	if y.needsQuoting(s) {
		buf.WriteByte('"')
		writeYAMLEscaped(buf, s)
		buf.WriteByte('"')

		return
	}

	buf.WriteString(s)
}

// needsQuoting reports whether a plain YAML scalar would be misinterpreted and thus needs double-quoting.
// It is intentionally less aggressive than containsSpecialChars so that ordinary messages (which often
// contain spaces) are kept unquoted.
func (y *YamlMarshaler) needsQuoting(s string) bool {
	if s == "" {
		return true
	}

	switch s[0] {
	case ' ', '!', '&', '*', '?', '|', '>', '%', '@', '`', '"', '\'', '#', ',', '[', ']', '{', '}', ':', '-':
		return true
	}

	if s[len(s)-1] == ' ' {
		return true
	}

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch {
		case c < 0x20 || c == '"' || c == '\\':
			return true
		case c == ':' && (i+1 == len(s) || s[i+1] == ' '):
			return true
		case c == '#' && i > 0 && s[i-1] == ' ':
			return true
		}
	}

	return false
}

// writeYAMLEscaped writes s into the buffer escaping the characters that would break a
// double-quoted YAML scalar (double quotes, backslashes and control chars).
func writeYAMLEscaped(buf *bytes.Buffer, s string) {
	start := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != '"' && c != '\\' && c >= 0x20 {
			continue
		}

		if start < i {
			buf.WriteString(s[start:i])
		}

		switch c {
		case '"':
			buf.WriteString(`\"`)
		case '\\':
			buf.WriteString(`\\`)
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		default:
			buf.WriteString(`\x`)
			buf.WriteByte(hexDigits[c>>4])
			buf.WriteByte(hexDigits[c&0x0F])
		}

		start = i + 1
	}

	if start < len(s) {
		buf.WriteString(s[start:])
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
