package encoders

import (
	"bytes"
	"errors"
	s "github.com/pho3b/tiny-logger/shared"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildMsg(t *testing.T) {
	encoder := &BaseEncoder{}

	// Test with multiple arguments
	result := encoder.castAndConcatenate("This", "is", "a", "test")
	assert.Equal(t, "This is a test", result)

	// Test with a single argument
	result = encoder.castAndConcatenate("SingleArgument")
	assert.Equal(t, "SingleArgument", result)

	// Test with no arguments
	result = encoder.castAndConcatenate()
	assert.Equal(t, "", result)

	// Test with mixed data types
	result = encoder.castAndConcatenate("Mixed", 123, true, 45.6)
	assert.Equal(t, "Mixed 123 true 45.6", result)

	// Test with rune and int64 types and struct
	result = encoder.castAndConcatenate('A', int64(43), errors.New("my error"))
	assert.Equal(t, "A 43 my error", result)
}

func TestAreAllNil(t *testing.T) {
	encoder := &BaseEncoder{}

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

	baseEncoder := &BaseEncoder{}
	assert.Equal(t, s.EncoderType(""), baseEncoder.GetType())
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
