package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrintDateTime_AddDateOnly(t *testing.T) {
	printer := DateTimePrinter{}
	mockDate := time.Now().Format("02/01/2006")

	result := printer.PrintDateTime(true, false)

	expected := fmt.Sprintf("[%s]", mockDate)
	assert.Equal(t, expected, result, "Expected result to include only the date in [DD/MM/YYYY] format")
}

func TestPrintDateTime_AddTimeOnly(t *testing.T) {
	printer := DateTimePrinter{}
	mockT := time.Now().Format("15:04:05")

	result := printer.PrintDateTime(false, true)

	expected := fmt.Sprintf("[%s]", mockT)
	assert.Equal(t, expected, result, "Expected result to include only the time in [HH:MM:SS] format")
}

func TestPrintDateTime_AddDateAndTime(t *testing.T) {
	printer := DateTimePrinter{}
	mockDate := time.Now().Format("02/01/2006")
	mockT := time.Now().Format("15:04:05")

	result := printer.PrintDateTime(true, true)

	expected := fmt.Sprintf("[%s %s]", mockDate, mockT)
	assert.Equal(t, expected, result, "Expected result to include both date and time in [DD/MM/YYYY HH:MM:SS] format")
}

func TestPrintDateTime_NoDateNoTime(t *testing.T) {
	printer := DateTimePrinter{}
	result := printer.PrintDateTime(false, false)

	assert.Equal(t, "", result, "Expected result to be an empty string when both addDate and addTime are false")
}
