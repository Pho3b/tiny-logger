package services

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	jsonCharOverhead = 80
	averageExtraLen  = 30
	// firstPrintableASCII is the first printable ASCII byte; bytes below it are
	// control characters that must be escaped in JSON output.
	firstPrintableASCII = 0x20
)

// hexDigits is the lookup table used to render control characters as hex escapes.
// It is shared by the JSON and YAML escapers (writeJSONEscaped and writeYAMLEscaped).
const hexDigits = "0123456789abcdef"

// JsonLogEntry represents a structured log entry that can be marshaled to JSON format.
// All fields except Message are optional and will be omitted if empty.
type JsonLogEntry struct {
	Level   string `json:"level,omitempty"`
	Date    string `json:"date,omitempty"`
	Time    string `json:"time,omitempty"`
	Message string `json:"msg"`
	UnixTS  string `json:"ts,omitempty"`
	Extras  []any  `json:"extras,omitempty"`
}

// JsonMarshaler provides custom JSON marshaling functionality optimized for log entries.
type JsonMarshaler struct {
}

// MarshalInto converts a JsonLogEntry into a JSON-formatted byte slice and adds it to the given buffer
// to minimize allocations during marshaling.
func (j *JsonMarshaler) MarshalInto(buf *bytes.Buffer, logEntry JsonLogEntry) {
	extrasLen := len(logEntry.Extras)
	buf.Grow(jsonCharOverhead + (averageExtraLen * extrasLen))

	buf.WriteByte('{')
	j.writeLogEntryProperties(buf, logEntry.Level, logEntry.Date, logEntry.Time, logEntry.UnixTS)

	buf.WriteString("\"msg\":\"")
	writeJSONEscaped(buf, logEntry.Message)
	buf.WriteByte('"')

	if extrasLen > 0 {
		buf.WriteString(",\"extras\":{")

		for i := 0; i < extrasLen; i += 2 {
			if i < extrasLen {
				buf.WriteByte('"')
				j.writeValue(buf, logEntry.Extras[i], true)
				buf.WriteString(`":`)

				k := i + 1
				if k < extrasLen {
					j.writeValue(buf, logEntry.Extras[k], false)

					if k < extrasLen-1 {
						buf.WriteByte(',')
					}
				}
			}
		}

		if extrasLen%2 != 0 {
			buf.WriteString("null")
		}

		buf.WriteByte('}')
	}

	buf.WriteByte('}')
}

// writeValue writes a value to the buffer with appropriate JSON formatting.
// The method handles different types (string, rune, int, int64, float64, bool)
// with special consideration for whether the value is being written as a key or value.
func (j *JsonMarshaler) writeValue(buf *bytes.Buffer, v any, isKey bool) {
	switch val := v.(type) {
	case string:
		if isKey {
			writeJSONEscaped(buf, val)
		} else {
			buf.WriteByte('"')
			writeJSONEscaped(buf, val)
			buf.WriteByte('"')
		}
	case rune:
		if isKey {
			writeJSONEscaped(buf, string(val))
		} else {
			buf.WriteByte('"')
			writeJSONEscaped(buf, string(val))
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
		if isKey {
			writeJSONEscaped(buf, fmt.Sprint(val))
		} else {
			buf.WriteByte('"')
			writeJSONEscaped(buf, fmt.Sprint(val))
			buf.WriteByte('"')
		}
	}
}

// writeLogEntryProperties writes the standard log entry properties to the buffer.
// Only non-empty properties are written, each followed by a comma.
func (j *JsonMarshaler) writeLogEntryProperties(
	buf *bytes.Buffer,
	level string,
	date string,
	time string,
	unixTS string,
) {
	if level != "" {
		buf.WriteString("\"level\":\"")
		buf.WriteString(level)
		buf.WriteByte('"')
		buf.WriteByte(',')
	}

	if unixTS != "" {
		buf.WriteString("\"ts\":\"")
		buf.WriteString(unixTS)
		buf.WriteByte('"')
		buf.WriteByte(',')

		return
	}

	if date != "" && time != "" {
		buf.WriteString("\"datetime\":\"")
		buf.WriteString(date)
		buf.WriteByte(' ')
		buf.WriteString(time)
		buf.WriteByte('"')
		buf.WriteByte(',')

		return
	}

	if date != "" {
		buf.WriteString("\"date\":\"")
		buf.WriteString(date)
		buf.WriteByte('"')
		buf.WriteByte(',')
	}

	if time != "" {
		buf.WriteString("\"time\":\"")
		buf.WriteString(time)
		buf.WriteByte('"')
		buf.WriteByte(',')
	}
}

// writeJSONEscaped writes s into the buffer escaping characters that would produce
// invalid JSON: double quotes, backslashes, and control chars below 0x20.
// It keeps a fast path that writes the whole string in one shot when no escaping is needed.
func writeJSONEscaped(buf *bytes.Buffer, s string) {
	start := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != '"' && c != '\\' && c >= firstPrintableASCII {
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
			buf.WriteString(`\u00`)
			buf.WriteByte(hexDigits[c>>4])
			buf.WriteByte(hexDigits[c&0x0F])
		}

		start = i + 1
	}

	if start < len(s) {
		buf.WriteString(s[start:])
	}
}

func NewJsonMarshaler() JsonMarshaler {
	return JsonMarshaler{}
}
