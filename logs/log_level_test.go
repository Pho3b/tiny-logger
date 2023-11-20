package logs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
