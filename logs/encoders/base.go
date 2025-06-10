package encoders

import (
	"bytes"
	"fmt"
	"github.com/pho3b/tiny-logger/shared"
	"strconv"
)

const averageWordLen = 30

type BaseEncoder struct {
	encoderType shared.EncoderType
}

// concatenate returns a string containing all the given arguments cast to strings concatenated with a white space.
func (b *BaseEncoder) concatenate(args ...interface{}) string {
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
func (b *BaseEncoder) areAllNil(args ...interface{}) bool {
	for _, arg := range args {
		if arg != nil {
			return false
		}
	}

	return true
}

func (b *BaseEncoder) GetType() shared.EncoderType {
	return b.encoderType
}
