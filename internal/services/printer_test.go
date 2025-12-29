package services

import (
	"testing"

	c "github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/stretchr/testify/assert"
)

func TestPrintColors_EnableColorsTrue(t *testing.T) {
	printer := Printer{}

	result := printer.RetrieveColorsFromLogLevel(true, log_level.FatalErrorLvl)
	assert.Equal(t, c.Magenta, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")

	result = printer.RetrieveColorsFromLogLevel(true, log_level.ErrorLvl)
	assert.Equal(t, c.Red, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")

	result = printer.RetrieveColorsFromLogLevel(true, log_level.WarnLvl)
	assert.Equal(t, c.Yellow, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")

	result = printer.RetrieveColorsFromLogLevel(true, log_level.InfoLvl)
	assert.Equal(t, c.Cyan, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")

	result = printer.RetrieveColorsFromLogLevel(true, log_level.DebugLvl)
	assert.Equal(t, c.Gray, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")
}

func TestPrintColors_EnableColorsFalse(t *testing.T) {
	printer := Printer{}

	result := printer.RetrieveColorsFromLogLevel(false, log_level.DebugLvl)
	assert.Equal(t, c.Color(""), result[0], "Expected first element to be an empty string when colors are disabled")
	assert.Equal(t, c.Color(""), result[1], "Expected second element to be an empty string when colors are disabled")

	result = printer.RetrieveColorsFromLogLevel(false, log_level.InfoLvl)
	assert.Equal(t, c.Color(""), result[0], "Expected first element to be an empty string when colors are disabled")
	assert.Equal(t, c.Color(""), result[1], "Expected second element to be an empty string when colors are disabled")
}
