package encoders

import (
	"bytes"
	"errors"
	"os"
	"sync"
	"testing"

	s "github.com/pho3b/tiny-logger/shared"
	"github.com/stretchr/testify/assert"
)

func TestBuildMsg(t *testing.T) {
	encoder := newBaseEncoder()

	// Test with multiple arguments
	result := encoder.castToString("This")
	assert.Equal(t, "This", result)

	// Test with a single argument
	result = encoder.castToString("SingleArgument")
	assert.Equal(t, "SingleArgument", result)

	// Test with various argument types
	result = encoder.castToString(2)
	assert.Equal(t, "2", result)
	result = encoder.castToString(2.3)
	assert.Equal(t, "2.3", result)
	result = encoder.castToString(true)
	assert.Equal(t, "true", result)
	result = encoder.castToString(nil)
	assert.Equal(t, "<nil>", result)

	// Test with no arguments
	result = encoder.castToString("")
	assert.Equal(t, "", result)

	// Test with rune and int64 types and struct
	result = encoder.castToString(int64(43))
	assert.Equal(t, "43", result)
	result = encoder.castToString(int32(32234))
	assert.Equal(t, "緪", result)
	result = encoder.castToString(int8(3))
	assert.Equal(t, "3", result)

	result = encoder.castToString(errors.New("my error"))
	assert.Equal(t, "my error", result)

	result = encoder.castToString(struct{ test string }{"test"})
	assert.Equal(t, "{test}", result)
}

func TestBuildMsgWithCastAndConcatenateInto(t *testing.T) {
	encoder := newBaseEncoder()
	buf := &bytes.Buffer{}

	// Test with multiple arguments
	encoder.castAndConcatenateInto(buf, "This", "is", 'a', "test")
	assert.Equal(t, "This is a test", buf.String())
	buf.Reset()

	// Test with a single argument
	encoder.castAndConcatenateInto(buf, "SingleArgument")
	assert.Equal(t, "SingleArgument", buf.String())
	buf.Reset()

	// Test with various argument types
	encoder.castAndConcatenateInto(buf, "str", '\n', 2, 2.3, true, nil)
	assert.Equal(t, "str \n 2 2.3 true <nil>", buf.String())
	buf.Reset()

	// Test with no arguments
	encoder.castAndConcatenateInto(buf)
	assert.Equal(t, "", buf.String())
	buf.Reset()

	// Test with mixed data types
	encoder.castAndConcatenateInto(buf, "Mixed", 123, true, 45.6)
	assert.Equal(t, "Mixed 123 true 45.6", buf.String())
	buf.Reset()

	// Test with rune and int64 types and struct
	encoder.castAndConcatenateInto(buf, 'A', int64(43), errors.New("my error"))
	assert.Equal(t, "A 43 my error", buf.String())
	buf.Reset()
}

func TestAreAllNil(t *testing.T) {
	encoder := newBaseEncoder()

	// Test with all nil arguments
	result := encoder.areAllNil(nil, nil, nil)
	assert.True(t, result)

	// Test with no arguments
	result = encoder.areAllNil()
	assert.True(t, result)

	// Test with some nil and some non-nil arguments
	result = encoder.areAllNil(nil, "test", nil)
	assert.False(t, result)

	// Test with all non-nil arguments
	result = encoder.areAllNil("test", 123, true)
	assert.False(t, result)
}

func TestBaseEncoder_GetType(t *testing.T) {
	encoder := NewDefaultEncoder()
	assert.Equal(t, s.DefaultEncoderType, encoder.GetType())

	jsonEncoder := NewJSONEncoder()
	assert.Equal(t, s.JsonEncoderType, jsonEncoder.GetType())

	yamlEncoder := NewYAMLEncoder()
	assert.Equal(t, s.YamlEncoderType, yamlEncoder.GetType())

	baseEncoder := newBaseEncoder()
	assert.Equal(t, s.EncoderType(""), baseEncoder.GetType())
}

// newBaseEncoder initializes and returns a new baseEncoder instance with initialized bufferSyncPool.
func newBaseEncoder() *baseEncoder {
	encoder := &baseEncoder{}
	encoder.bufferSyncPool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	return encoder
}

// captureOutput redirects os.Stdout to capture the output of the function f
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	defer r.Close()

	origStdout := os.Stdout
	os.Stdout = w

	f()
	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

// captureErrorOutput redirects os.Stderr to capture the output of the function f
func captureErrorOutput(f func()) string {
	r, w, _ := os.Pipe()
	defer r.Close()

	origStderr := os.Stderr
	os.Stderr = w

	f()
	w.Close()
	os.Stderr = origStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
