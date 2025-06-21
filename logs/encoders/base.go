package encoders

import (
	"bytes"
	"fmt"
	s "github.com/pho3b/tiny-logger/shared"
	"os"
	"strconv"
)

const averageWordLen = 30

type BaseEncoder struct {
	encoderType s.EncoderType
}

// castAndConcatenate returns a string containing all the given arguments cast to string and concatenated by a white space.
func (b *BaseEncoder) castAndConcatenate(args ...any) string {
	var res bytes.Buffer
	res.Grow(averageWordLen * len(args)) // Assuming an average word length of 30 chars

	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			res.WriteString(v)
		case rune:
			res.WriteRune(v)
		case int:
			res.WriteString(strconv.Itoa(v))
		case int64:
			res.WriteString(strconv.FormatInt(v, 10))
		case float64:
			res.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			res.WriteString(strconv.FormatBool(v))
		default:
			// Using the slower fmt.Sprint only for unknown types
			res.WriteString(fmt.Sprint(v))
		}

		if i < len(args)-1 {
			res.WriteByte(' ')
		}
	}

	return res.String()
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
