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
		dateTimePrinter.UpdateDateTimeFormat(shared.IT)
		dateRes, timeRes, unixTs := dateTimePrinter.RetrieveDateTime(true, true)
		assert.Empty(t, unixTs)
		assert.NotEmpty(t, dateRes)
		assert.NotEmpty(t, timeRes)
		assert.Equal(t, "01/11/2023 15:30:45", dateRes+" "+timeRes)
	})

	t.Run("Return date only", func(t *testing.T) {
		dateTimePrinter.UpdateDateTimeFormat(shared.IT)
		dateRes, timeRes, unixTs := dateTimePrinter.RetrieveDateTime(true, false)
		assert.Equal(t, "01/11/2023", dateRes)
		assert.Equal(t, "", timeRes)
		assert.Equal(t, "", unixTs)
	})

	t.Run("Return time only", func(t *testing.T) {
		dateTimePrinter.UpdateDateTimeFormat(shared.IT)
		dateRes, timeRes, unixTs := dateTimePrinter.RetrieveDateTime(false, true)
		assert.Equal(t, "", dateRes)
		assert.Equal(t, "15:30:45", timeRes)
		assert.Equal(t, "", unixTs)
	})

	t.Run("Return empty string when both flags are false", func(t *testing.T) {
		dateTimePrinter.UpdateDateTimeFormat(shared.IT)
		dateRes, timeRes, dateTimeRes := dateTimePrinter.RetrieveDateTime(false, false)
		assert.Equal(t, "", dateRes)
		assert.Equal(t, "", timeRes)
		assert.Equal(t, "", dateTimeRes)
	})
}

func TestDateTimePrinter_Formats(t *testing.T) {
	// Using a future time to prevent the background goroutines from spinning
	// (since time.Sleep has depended on a real clock vs. mocked time).
	// 2099-11-01 15:30:45 UTC
	fixedFutureTime := time.Date(2099, time.November, 1, 15, 30, 45, 0, time.UTC)

	dateTimePrinter := &DateTimePrinter{
		timeNow: func() time.Time {
			return fixedFutureTime
		},
	}

	tests := []struct {
		name         string
		format       shared.DateTimeFormat
		wantDate     string
		wantTime     string
		wantCombined string
	}{
		{
			name:         "IT Format",
			format:       shared.IT,
			wantDate:     "01/11/2099",
			wantTime:     "15:30:45",
			wantCombined: "01/11/2099 15:30:45",
		},
		{
			name:         "US Format",
			format:       shared.US,
			wantDate:     "11/01/2099",
			wantTime:     "03:30:45 PM",
			wantCombined: "11/01/2099 03:30:45 PM",
		},
		{
			name:         "JP Format",
			format:       shared.JP,
			wantDate:     "2099/11/01",
			wantTime:     "03:30:45 PM",
			wantCombined: "2099/11/01 03:30:45 PM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dateTimePrinter.UpdateDateTimeFormat(tt.format)
			dateRes, timeRes, unixTs := dateTimePrinter.RetrieveDateTime(true, true)
			assert.NotEmpty(t, dateRes)
			assert.NotEmpty(t, timeRes)
			assert.Empty(t, unixTs)
			assert.Equal(t, tt.wantCombined, dateRes+" "+timeRes)

			// Also test individual components
			d, tRes, _ := dateTimePrinter.RetrieveDateTime(true, false)
			assert.Equal(t, tt.wantDate, d)
			assert.Empty(t, tRes)

			d, tRes, _ = dateTimePrinter.RetrieveDateTime(false, true)
			assert.Empty(t, d)
			assert.Equal(t, tt.wantTime, tRes)
		})
	}
}

func TestDateTimePrinter_UnixTimestamp(t *testing.T) {
	// Use future time to avoid spin loop
	fixedFutureTime := time.Date(2099, time.November, 1, 15, 30, 45, 0, time.UTC)
	expectedUnix := "4097230245" // 2099-11-01 15:30:45 UTC

	dateTimePrinter := &DateTimePrinter{
		timeNow: func() time.Time {
			return fixedFutureTime
		},
	}

	dateTimePrinter.UpdateDateTimeFormat(shared.UnixTimestamp)

	t.Run("Return unix timestamp combined", func(t *testing.T) {
		dateRes, timeRes, dateTimeRes := dateTimePrinter.RetrieveDateTime(true, true)
		assert.Empty(t, dateRes)
		assert.Empty(t, timeRes)
		assert.Equal(t, expectedUnix, dateTimeRes)
	})

	t.Run("Return unix timestamp with date flag only", func(t *testing.T) {
		dateRes, timeRes, dateTimeRes := dateTimePrinter.RetrieveDateTime(true, false)
		assert.Empty(t, dateRes)
		assert.Empty(t, timeRes)
		assert.Equal(t, expectedUnix, dateTimeRes)
	})

	t.Run("Return unix timestamp with time flag only", func(t *testing.T) {
		dateRes, timeRes, dateTimeRes := dateTimePrinter.RetrieveDateTime(false, true)
		assert.Empty(t, dateRes)
		assert.Empty(t, timeRes)
		assert.Equal(t, expectedUnix, dateTimeRes)
	})
}

func TestNewDateTimePrinter(t *testing.T) {
	assert.NotNil(t, NewDateTimePrinter())
	assert.IsType(t, &DateTimePrinter{}, NewDateTimePrinter())
}

func TestDateTimePrinter_FullSecondUpdate(t *testing.T) {
	dateTimePrinter := NewDateTimePrinter()

	t.Run("Return both date and time", func(t *testing.T) {
		dateTimePrinter.UpdateDateTimeFormat(shared.IT)
		dateRes, timeRes, _ := dateTimePrinter.RetrieveDateTime(true, true)
		assert.NotEmpty(t, dateRes)
		assert.NotEmpty(t, timeRes)

		prevTime := dateTimePrinter.currentTime.Load()
		time.Sleep(2 * time.Second)
		currTime := dateTimePrinter.currentTime.Load()

		assert.NotEqual(t, prevTime, currTime,
			"The current %s time should have changed from previous time", currTime, prevTime)
	})
}
