package services

import (
	"fmt"
	"time"
)

type DateTimePrinter struct {
}

// PrintDateTime returns a formatted date and/or time string based on the provided flags.
// If addDate is true, the current date is included in "DD/MM/YYYY" format.
// If addTime is true, the current time is included in "HH:MM:SS" format.
// If neither addDate nor addTime is true, an empty string is returned.
func (d *DateTimePrinter) PrintDateTime(addDate bool, addTime bool) string {
	if !addDate && !addTime {
		return ""
	}

	dateRes := ""
	timeRes := ""

	if addDate {
		dateRes = fmt.Sprintf("%s", time.Now().Format("02/01/2006"))
	}

	if addTime {
		whiteSpace := " "

		if !addDate {
			whiteSpace = ""
		}

		timeRes = fmt.Sprintf("%s%s", whiteSpace, time.Now().Format("15:04:05"))
	}

	return fmt.Sprintf("[%s%s]", dateRes, timeRes)
}
