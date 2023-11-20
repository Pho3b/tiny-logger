package logs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAreAllNil(t *testing.T) {
	assert.True(t, areAllNil(nil, nil, nil, nil))
	assert.False(t, areAllNil(nil, 3, nil, nil))
	assert.True(t, areAllNil())
	assert.False(t, areAllNil("test", 43, nil))
}

func TestRetrieveLogLvlFromEnv(t *testing.T) {
	assert.Equal(t, DebugLvl, retrieveLogLvlFromEnv(""))

	testEnvVar := "MY_INSTANCE_LOGS_LVL_2"
	_ = os.Setenv(testEnvVar, InfoLvlName)
	assert.Equal(t, InfoLvl, retrieveLogLvlFromEnv(testEnvVar))
	assert.NotEqual(t, DebugLvl, retrieveLogLvlFromEnv(testEnvVar))
}

func TestRetrieveLogLvlIntFromName(t *testing.T) {
	assert.Equal(t, InfoLvl, retrieveLogLvlIntFromName(InfoLvlName))
	assert.Equal(t, DebugLvl, retrieveLogLvlIntFromName(DebugLvlName))
	assert.Equal(t, WarnLvl, retrieveLogLvlIntFromName(WarnLvlName))
	assert.Equal(t, ErrorLvl, retrieveLogLvlIntFromName(ErrorLvlName))
	assert.Equal(t, DebugLvl, retrieveLogLvlIntFromName("non-existing-lvl-name"))
	assert.NotEqual(t, ErrorLvl, retrieveLogLvlIntFromName(""))
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
