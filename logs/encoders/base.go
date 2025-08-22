package encoders

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"sync"

	s "github.com/pho3b/tiny-logger/shared"
)

const averageWordLen = 35

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

type BaseEncoder struct {
	encoderType s.EncoderType
}

// castAndConcatenate returns a string containing all the given arguments cast to string and concatenated by a white space.
func (b *BaseEncoder) castAndConcatenate(args ...any) string {
	argsLen := len(args)
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	buf.Grow(averageWordLen * argsLen)

	for i, arg := range args {
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

		if i < argsLen-1 {
			buf.WriteByte(' ')
		}
	}

	res := buf.String()
	bufferPool.Put(buf)

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
func (b *BaseEncoder) printLog(outType s.OutputType, msgBuffer bytes.Buffer, newLine bool) {
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

func (b *BaseEncoder) GetType() s.EncoderType {
	return b.encoderType
}
