package encoders

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildMsg(t *testing.T) {
	encoder := &BaseEncoder{}

	// Test with multiple arguments
	result := encoder.buildMsg("This", "is", "a", "test")
	assert.Equal(t, "This is a test", result)

	// Test with a single argument
	result = encoder.buildMsg("SingleArgument")
	assert.Equal(t, "SingleArgument", result)

	// Test with no arguments
	result = encoder.buildMsg()
	assert.Equal(t, "", result)

	// Test with mixed data types
	result = encoder.buildMsg("Mixed", 123, true, 45.6)
	assert.Equal(t, "Mixed 123 true 45.6", result)
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
