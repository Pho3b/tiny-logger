package services

import (
	"bytes"
	"fmt"
	"strconv"
)

// JsonLogEntry represents a structured log entry that can be marshaled to JSON format.
// All fields except Message are optional and will be omitted if empty.
type JsonLogEntry struct {
	Level    string `json:"level,omitempty"`
	Date     string `json:"date,omitempty"`
	Time     string `json:"time,omitempty"`
	DateTime string `json:"datetime,omitempty"`
	Message  string `json:"msg"`
	Extras   []any  `json:"extras,omitempty"`
}

// JsonMarshaler provides custom JSON marshaling functionality optimized for log entries.
type JsonMarshaler struct {
}

// Marshal converts a JsonLogEntry into a JSON-formatted byte slice.
// It uses a buffer-based approach to minimize allocations during marshaling.
func (j *JsonMarshaler) Marshal(logEntry JsonLogEntry) []byte {
	var res bytes.Buffer
	extrasLen := len(logEntry.Extras)
	res.Grow(jsonCharOverhead + (averageExtraLen * extrasLen))

	res.WriteByte('{')
	j.writeLogEntryProperties(&res, logEntry.Level, logEntry.Date, logEntry.Time, logEntry.DateTime)

	res.WriteString("\"msg\":\"")
	res.WriteString(logEntry.Message)
	res.WriteByte('"')

	if extrasLen > 0 {
		res.WriteString(",\"extras\":{")

		for i := 0; i < extrasLen; i += 2 {
			if i < extrasLen {
				res.WriteByte('"')
				j.writeValue(&res, logEntry.Extras[i], true)
				res.WriteString(`":`)

				k := i + 1
				if k < extrasLen {
					j.writeValue(&res, logEntry.Extras[k], false)

					if k < extrasLen-1 {
						res.WriteByte(',')
					}
				}
			}
		}

		if extrasLen%2 != 0 {
			res.WriteString("null")
		}

		res.WriteByte('}')
	}

	res.WriteByte('}')
	return res.Bytes()
}

func (j *JsonMarshaler) MarshalInto(buf *bytes.Buffer, logEntry JsonLogEntry) {
	extrasLen := len(logEntry.Extras)
	buf.Grow(jsonCharOverhead + (averageExtraLen * extrasLen))

	buf.WriteByte('{')
	j.writeLogEntryProperties(buf, logEntry.Level, logEntry.Date, logEntry.Time, logEntry.DateTime)

	buf.WriteString("\"msg\":\"")
	buf.WriteString(logEntry.Message)
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
			buf.WriteString(val)
		} else {
			buf.WriteByte('"')
			buf.WriteString(val)
			buf.WriteByte('"')
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
		if isKey {
			buf.WriteString(fmt.Sprint(val))
		} else {
			buf.WriteByte('"')
			buf.WriteString(fmt.Sprint(val))
			buf.WriteByte('"')
		}
	}
}

// writeLogEntryProperties writes the standard log entry properties to the buffer.
// Only non-empty properties are written, each followed by a comma.
func (j *JsonMarshaler) writeLogEntryProperties(buf *bytes.Buffer, level string, date string, time string, dateTime string) {
	if level != "" {
		buf.WriteString("\"level\":\"")
		buf.WriteString(level)
		buf.WriteByte('"')
		buf.WriteByte(',')
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

	if dateTime != "" {
		buf.WriteString("\"datetime\":\"")
		buf.WriteString(dateTime)
		buf.WriteByte('"')
		buf.WriteByte(',')
	}
}
