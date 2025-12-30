package encoders

import (
	"bytes"
	"fmt"
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
			buf.Write(strconv.AppendInt(buf.AvailableBuffer(), int64(v), 10))
		case int64:
			buf.Write(strconv.AppendInt(buf.AvailableBuffer(), v, 10))
		case float64:
			buf.Write(strconv.AppendFloat(buf.AvailableBuffer(), v, 'f', -1, 64))
		case bool:
			buf.Write(strconv.AppendBool(buf.AvailableBuffer(), v))
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
		// Using the slower fmt.Sprint only for unknown types
		return fmt.Sprint(v)
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
