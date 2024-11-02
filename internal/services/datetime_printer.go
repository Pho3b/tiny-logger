package services

import (
	"fmt"
	"time"
)

type DateTimePrinter struct {
}

func (d *DateTimePrinter) PrintDateTime(addDatetime bool) string {
	dateTime := ""

	if addDatetime {
		dateTime = fmt.Sprintf("[%s]", time.Now().Format("02/01/2006 15:04:05"))
	}

	return dateTime
}
