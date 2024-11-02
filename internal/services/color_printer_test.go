package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	c "gitlab.com/docebo/libraries/go/tiny-logger/logs/colors"
)

func TestPrintColors_EnableColorsTrue(t *testing.T) {
	printer := ColorsPrinter{}
	color := c.Cyan

	result := printer.PrintColors(true, color)

	assert.Equal(t, color, result[0], "Expected first element to be the provided color")
	assert.Equal(t, c.Reset, result[1], "Expected second element to be the reset color")
}

func TestPrintColors_EnableColorsFalse(t *testing.T) {
	printer := ColorsPrinter{}
	color := c.Color("blue")

	result := printer.PrintColors(false, color)

	assert.Equal(t, c.Color(""), result[0], "Expected first element to be an empty string when colors are disabled")
	assert.Equal(t, c.Color(""), result[1], "Expected second element to be an empty string when colors are disabled")
}
