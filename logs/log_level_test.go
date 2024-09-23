package logs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRetrieveLogLvlFromEnv(t *testing.T) {
	assert.Equal(t, DebugLvl, retrieveLogLvlFromEnv(""))

	testEnvVar := "MY_INSTANCE_LOGS_LVL_2"
	_ = os.Setenv(testEnvVar, string(InfoLvlName))
	assert.Equal(t, InfoLvl, retrieveLogLvlFromEnv(testEnvVar))
	assert.NotEqual(t, DebugLvl, retrieveLogLvlFromEnv(testEnvVar))
}

func TestLogLevel_LvlName(t *testing.T) {
	logLvl := LogLevel{
		lvl:         2,
		envVariable: "test-env-var",
	}
	assert.Equal(t, InfoLvlName, logLvl.LvlName())

	logLvl.lvl = 3
	assert.Equal(t, DebugLvlName, logLvl.LvlName())
}

func TestLogLvlIntValue(t *testing.T) {
	logLvl := LogLevel{
		lvl:         2,
		envVariable: "test-env-var",
	}
	assert.Equal(t, InfoLvl, logLvl.LvlIntValue())

	logLvl.lvl = 3
	assert.Equal(t, DebugLvl, logLvl.LvlIntValue())
}
