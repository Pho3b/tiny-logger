package encoders

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"sync"

	s "github.com/pho3b/tiny-logger/shared"
)

const (
	averageWordLen      = 30
	defaultCharOverhead = 50
)

type baseEncoder struct {
	encoderType    s.EncoderType
	bufferSyncPool sync.Pool
}

// castAndConcatenateInto writes all the given arguments cast to string and concatenated by a white space into the given buffer.
func (b *baseEncoder) castAndConcatenateInto(buf *bytes.Buffer, args ...any) {
	argsLen := len(args)
	buf.Grow(averageWordLen * argsLen)

	for i, arg := range args {
		if i > 0 {
			buf.WriteByte(' ')
		}

		switch v := arg.(type) {
		case string:
			buf.WriteString(v)
		case rune:
			buf.WriteRune(v)
		case int:
			buf.WriteString(strconv.Itoa(v))
		case int64:
			buf.WriteString(strconv.FormatInt(v, 10))
		case float64:
			buf.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			buf.WriteString(strconv.FormatBool(v))
		case fmt.Stringer:
			buf.WriteString(v.String())
		case error:
			buf.WriteString(v.Error())
		default:
			// Using the slower fmt.Sprint only for unknown types
			buf.WriteString(fmt.Sprint(v))
		}
	}
}

// castToString is a fast casting method that returns the given argument as a string.
func (b *baseEncoder) castToString(arg any) string {
	switch v := arg.(type) {
	case string:
		return v
	case rune:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case fmt.Stringer:
		return v.String()
	case error:
		return v.Error()
	default:
		return fmt.Sprint(v)
	}
}

// areAllNil returns true if all the given args are 'nil', false otherwise.
func (b *baseEncoder) areAllNil(args ...any) bool {
	for _, arg := range args {
		if arg != nil {
			return false
		}
	}

	return true
}

// printLog prints the given msgBuffer to the given outputType (stdout or stderr).
// If 'file' is not nil, the message is written to the file.
// If 'newLine' is true, a new line is added at the end of the message.
func (b *baseEncoder) printLog(outType s.OutputType, msgBuffer *bytes.Buffer, newLine bool, file *os.File) {
	if newLine {
		msgBuffer.WriteByte('\n')
	}

	if file != nil {
		_, err := file.Write(msgBuffer.Bytes())
		if err != nil {
			_, _ = os.Stderr.Write(msgBuffer.Bytes())
		}

		return
	}

	switch outType {
	case s.StdOutput:
		_, _ = os.Stdout.Write(msgBuffer.Bytes())
	case s.StdErrOutput:
		_, _ = os.Stderr.Write(msgBuffer.Bytes())
	}
}

// getBuffer returns a new bytes buffer from the pool.
// If the pool is empty, a new buffer is created.
func (b *baseEncoder) getBuffer() *bytes.Buffer {
	return b.bufferSyncPool.Get().(*bytes.Buffer)
}

// putBuffer puts the given bytes buffer back to the pool.
func (b *baseEncoder) putBuffer(buf *bytes.Buffer) {
	buf.Reset()
	b.bufferSyncPool.Put(buf)
}

// GetType returns the encoder type.
func (b *baseEncoder) GetType() s.EncoderType {
	return b.encoderType
}
