package colors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsColorValid(t *testing.T) {
	assert.True(t, IsColorValid(Red))
	assert.True(t, IsColorValid(Magenta))
	assert.True(t, IsColorValid(Yellow))
	assert.True(t, IsColorValid(Cyan))
	assert.True(t, IsColorValid(Gray))
	assert.True(t, IsColorValid(White))
	assert.True(t, IsColorValid(Black))
	assert.False(t, IsColorValid("  "))
	assert.False(t, IsColorValid("lfdjlfkasd"))
	assert.False(t, IsColorValid("\\testno"))
}
