package services

import (
	"testing"
	"time"

	"github.com/pho3b/tiny-logger/shared"
	"github.com/stretchr/testify/assert"
)

func TestDateTimePrinter_PrintDateTime(t *testing.T) {
	// Create a DateTimePrinter instance with a mocked timeNow function.
	dateTimePrinter := &DateTimePrinter{
		timeNow: func() time.Time { // mock time.Now to return a fixed time
			return time.Date(2023, time.November, 1, 15, 30, 45, 0, time.UTC)
		},
	}

	t.Run("Return both date and time", func(t *testing.T) {
		dateRes, timeRes, dateTimeRes := dateTimePrinter.RetrieveDateTime(true, true, shared.IT)
		assert.Empty(t, dateRes)
		assert.Empty(t, timeRes)
		assert.Equal(t, "01/11/2023 15:30:45", dateTimeRes)
	})

	t.Run("Return date only", func(t *testing.T) {
		dateRes, timeRes, dateTimeRes := dateTimePrinter.RetrieveDateTime(true, false, shared.IT)
		assert.Equal(t, "01/11/2023", dateRes)
		assert.Equal(t, "", timeRes)
		assert.Equal(t, "", dateTimeRes)
	})

	t.Run("Return time only", func(t *testing.T) {
		dateRes, timeRes, dateTimeRes := dateTimePrinter.RetrieveDateTime(false, true, shared.IT)
		assert.Equal(t, "", dateRes)
		assert.Equal(t, "15:30:45", timeRes)
		assert.Equal(t, "", dateTimeRes)
	})

	t.Run("Return empty string when both flags are false", func(t *testing.T) {
		dateRes, timeRes, dateTimeRes := dateTimePrinter.RetrieveDateTime(false, false, shared.IT)
		assert.Equal(t, "", dateRes)
		assert.Equal(t, "", timeRes)
		assert.Equal(t, "", dateTimeRes)
	})
}

func TestNewDateTimePrinter(t *testing.T) {
	assert.NotNil(t, NewDateTimePrinter())
	assert.IsType(t, DateTimePrinter{}, NewDateTimePrinter())
}

func TestDateTimePrinter_FullSecondUpdate(t *testing.T) {
	dateTimePrinter := NewDateTimePrinter()

	t.Run("Return both date and time", func(t *testing.T) {
		dateRes, timeRes, _ := dateTimePrinter.RetrieveDateTime(true, true, shared.IT)
		assert.Empty(t, dateRes)
		assert.Empty(t, timeRes)

		prevTime := dateTimePrinter.currentTime.Load()
		time.Sleep(2 * time.Second)
		currTime := dateTimePrinter.currentTime.Load()

		assert.NotEqual(t, prevTime, currTime,
			"The current %s time should have changed from previous time", currTime, prevTime)
	})
}
