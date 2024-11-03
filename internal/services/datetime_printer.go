package services

import (
	"fmt"
	"time"
)

type DateTimePrinter struct {
	timeNow func() time.Time // Function to get current time, allows mocking for tests
}

// PrintDateTime returns formatted date and/or time strings based on input flags.
// - If addDate is true, the method returns the current date in "DD/MM/YYYY" format.
// - If addTime is true, it returns the current time in "HH:MM:SS" format.
// - If both addDate and addTime are true, both date and time are returned as separate strings with time prefixed by a space.
// - If neither addDate nor addTime is true, empty strings are returned.
func (d *DateTimePrinter) PrintDateTime(addDate bool, addTime bool) (dateRes string, timeRes string) {
	if addDate {
		dateRes = fmt.Sprintf("%s", d.timeNow().Format("02/01/2006"))
	}

	if addTime {
		whiteSpace := " "

		if !addDate {
			whiteSpace = ""
		}

		timeRes = fmt.Sprintf("%s%s", whiteSpace, d.timeNow().Format("15:04:05"))
	}

	return dateRes, timeRes
}

// NewDateTimePrinter initializes DateTimePrinter with default timeNow function.
func NewDateTimePrinter() *DateTimePrinter {
	return &DateTimePrinter{timeNow: time.Now}
}
