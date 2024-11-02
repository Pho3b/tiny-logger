package services

import (
	"fmt"
	"time"
)

type DateTimePrinter struct {
}

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
