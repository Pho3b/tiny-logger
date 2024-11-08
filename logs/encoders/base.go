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
	res := strings.Builder{}

	for _, arg := range args {
		res.WriteString(fmt.Sprintf("%v ", arg))
	}

	return strings.TrimSuffix(res.String(), " ")
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
