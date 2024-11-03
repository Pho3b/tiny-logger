package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateTimePrinter_PrintDateTime(t *testing.T) {
	// Create a DateTimePrinter instance with mocked timeNow function.
	dateTimePrinter := &DateTimePrinter{
		timeNow: func() time.Time { // mock time.Now to return a fixed time
			return time.Date(2023, time.November, 1, 15, 30, 45, 0, time.UTC)
		},
	}

	t.Run("Return both date and time", func(t *testing.T) {
		dateRes, timeRes := dateTimePrinter.PrintDateTime(true, true)
		assert.Equal(t, "01/11/2023", dateRes)
		assert.Equal(t, "15:30:45", timeRes)
	})

	t.Run("Return date only", func(t *testing.T) {
		dateRes, timeRes := dateTimePrinter.PrintDateTime(true, false)
		assert.Equal(t, "01/11/2023", dateRes)
		assert.Equal(t, "", timeRes)
	})

	t.Run("Return time only", func(t *testing.T) {
		dateRes, timeRes := dateTimePrinter.PrintDateTime(false, true)
		assert.Equal(t, "", dateRes)
		assert.Equal(t, "15:30:45", timeRes)
	})

	t.Run("Return empty string when both flags are false", func(t *testing.T) {
		dateRes, timeRes := dateTimePrinter.PrintDateTime(false, false)
		assert.Equal(t, "", dateRes)
		assert.Equal(t, "", timeRes)
	})
}
