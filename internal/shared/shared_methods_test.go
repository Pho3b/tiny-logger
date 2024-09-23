package shared

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
	"testing"
)

func TestAreAllNil(t *testing.T) {
	assert.True(t, areAllNil(nil, nil, nil, nil))
	assert.False(t, areAllNil(nil, 3, nil, nil))
	assert.True(t, areAllNil())
	assert.False(t, areAllNil("test", 43, nil))
}

func TestRetrieveLogLvlIntFromName(t *testing.T) {
	assert.Equal(t, log_level.InfoLvl, log_level.RetrieveLogLvlIntFromName(log_level.InfoLvlName))
	assert.Equal(t, log_level.DebugLvl, log_level.RetrieveLogLvlIntFromName(log_level.DebugLvlName))
	assert.Equal(t, log_level.WarnLvl, log_level.RetrieveLogLvlIntFromName(log_level.WarnLvlName))
	assert.Equal(t, log_level.ErrorLvl, log_level.RetrieveLogLvlIntFromName(log_level.ErrorLvlName))
	assert.Equal(t, log_level.DebugLvl, log_level.RetrieveLogLvlIntFromName("non-existing-Lvl-name"))
	assert.NotEqual(t, log_level.ErrorLvl, log_level.RetrieveLogLvlIntFromName(""))
}

func TestBuildMessage(t *testing.T) {
	msg := buildMsg("hello", "to", "you")
	assert.Equal(t, "hello to you", msg)

	msg = buildMsg("hello", "to", "you", 2)
	assert.Equal(t, "hello to you 2", msg)

	msg = buildMsg(nil)
	assert.Equal(t, "<nil>", msg)

	msg = buildMsg(nil, "")
	assert.Equal(t, "<nil> ", msg)
}
