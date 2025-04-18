package services

import (
	"time"
)

type DateTimePrinter struct {
	timeNow func() time.Time // Function to get current time, allows mocking for tests
}

// PrintDateTime returns formatted date and/or time strings based on input flags.
// - If addDate is true, the method returns the current date in "DD/MM/YYYY" format.
// - If addTime is true, it returns the current time in "HH:MM:SS" format.
// - If both addDate and addTime are true, dateTimeRes is returned as a unified string.
// - If neither addDate nor addTime is true, empty strings are returned.
func (d *DateTimePrinter) PrintDateTime(addDate bool, addTime bool) (dateRes string, timeRes string, dateTimeRes string) {
	now := d.timeNow()

	if addDate {
		dateRes = now.Format("02/01/2006")
	}

	if addTime {
		timeRes = now.Format("15:04:05")
	}

	if addDate && addTime {
		return "", "", dateRes + " " + timeRes
	}

	return dateRes, timeRes, dateTimeRes
}

// NewDateTimePrinter initializes DateTimePrinter with default timeNow function.
func NewDateTimePrinter() *DateTimePrinter {
	return &DateTimePrinter{timeNow: time.Now}
}
