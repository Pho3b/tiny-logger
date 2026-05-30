package services

import "bytes"

const (
	hexDigits           = "0123456789abcdef"
	firstPrintableASCII = 0x20
)

// WriteEscaped writes s into the buffer escaping characters that would produce
// invalid output: double quotes, backslashes, and control chars below 0x20.
// It keeps a fast path that writes the whole string in one shot when no escaping is needed.
func WriteEscaped(buf *bytes.Buffer, s string) {
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
