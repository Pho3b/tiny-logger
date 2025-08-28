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

type BaseEncoder struct {
	encoderType    s.EncoderType
	bufferSyncPool sync.Pool
}

// castAndConcatenateInto writes all the given arguments cast to string and concatenated by a white space into the given buffer.
func (b *BaseEncoder) castAndConcatenateInto(buf *bytes.Buffer, args ...any) {
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
		default:
			// Using the slower fmt.Sprint only for unknown types
			buf.WriteString(fmt.Sprint(v))
		}
	}
}

// castAndConcatenate returns a string containing all the given arguments cast to string and concatenated by a white space.
func (b *BaseEncoder) castAndConcatenate(args ...any) string {
	argsLen := len(args)
	buf := b.getBuffer()
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
		default:
			// Using the slower fmt.Sprint only for unknown types
			buf.WriteString(fmt.Sprint(v))
		}
	}

	res := buf.String()
	b.putBuffer(buf)

	return res
}

// areAllNil returns true if all the given args are 'nil', false otherwise.
func (b *BaseEncoder) areAllNil(args ...any) bool {
	for _, arg := range args {
		if arg != nil {
			return false
		}
	}

	return true
}

// printLog prints the given msgBuffer to the given outputType (stdout or stderr).
func (b *BaseEncoder) printLog(outType s.OutputType, msgBuffer *bytes.Buffer, newLine bool) {
	if newLine {
		msgBuffer.WriteByte('\n')
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
func (b *BaseEncoder) getBuffer() *bytes.Buffer {
	return b.bufferSyncPool.Get().(*bytes.Buffer)
}

// putBuffer puts the given bytes buffer back to the pool.
func (b *BaseEncoder) putBuffer(buf *bytes.Buffer) {
	buf.Reset()
	b.bufferSyncPool.Put(buf)
}

// GetType returns the encoder type.
func (b *BaseEncoder) GetType() s.EncoderType {
	return b.encoderType
}
