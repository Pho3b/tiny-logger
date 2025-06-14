package log_level

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRetrieveLogLvlFromEnv(t *testing.T) {
	assert.Equal(t, DebugLvl, RetrieveLogLvlFromEnv(""))

	testEnvVar := "MY_INSTANCE_LOGS_LVL_2"
	_ = os.Setenv(testEnvVar, string(InfoLvlName))
	assert.Equal(t, InfoLvl, RetrieveLogLvlFromEnv(testEnvVar))
	assert.NotEqual(t, DebugLvl, RetrieveLogLvlFromEnv(testEnvVar))
}

func TestLogLevel_LvlName(t *testing.T) {
	logLvl := LogLevel{
		Lvl:         2,
		EnvVariable: "test-env-var",
	}
	assert.Equal(t, InfoLvlName, logLvl.LvlName())

	logLvl.Lvl = 3
	assert.Equal(t, DebugLvlName, logLvl.LvlName())
}

func TestLogLvlIntValue(t *testing.T) {
	logLvl := LogLevel{
		Lvl:         2,
		EnvVariable: "test-env-var",
	}
	assert.Equal(t, InfoLvl, logLvl.LvlIntValue())

	logLvl.Lvl = 3
	assert.Equal(t, DebugLvl, logLvl.LvlIntValue())
}

func TestLogLvlName_String(t *testing.T) {
	logLvl := LogLevel{Lvl: 3}
	assert.Equal(t, DebugLvlName.String(), logLvl.LvlName().String())
	logLvl.Lvl = 2
	assert.Equal(t, InfoLvlName, logLvl.LvlName())
	logLvl.Lvl = 1
	assert.Equal(t, WarnLvlName.String(), logLvl.LvlName().String())
	logLvl.Lvl = 0
	assert.Equal(t, ErrorLvlName, logLvl.LvlName())
	logLvl.Lvl = -1
	assert.Equal(t, FatalErrorLvlName.String(), logLvl.LvlName().String())
}
