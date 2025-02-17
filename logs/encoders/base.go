package encoders

import (
	"fmt"
	"github.com/pho3b/tiny-logger/shared"
	"strings"
)

type BaseEncoder struct {
	encoderType shared.EncoderType
}

// buildMsg returns a string containing all the given arguments cast to strings concatenated with a white space.
func (b *BaseEncoder) buildMsg(args ...interface{}) string {
	var res strings.Builder
	res.Grow(30 * len(args)) // Assuming average word length of N chars

	for i, arg := range args {
		res.WriteString(fmt.Sprint(arg))

		if i < len(args)-1 {
			res.WriteString(" ")
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
