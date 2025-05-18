package services

import (
	"time"
)

type DateTimePrinter struct {
	timeNow     func() time.Time // Function to get current time, allows mocking for tests
	currentDate string
	currentTime string
}

// PrintDateTime returns formatted date and/or time strings based on input flags.
// - If addDate is true, the method returns the current date in "DD/MM/YYYY" format.
// - If addTime is true, it returns the current time in "HH:MM:SS" format.
// - If both addDate and addTime are true, dateTimeRes is returned as a unified string.
// - If neither addDate nor addTime is true, empty strings are returned.

// TODO: Add a cancel method for the spawned goroutines
func (d *DateTimePrinter) PrintDateTime(addDate bool, addTime bool) (dateRes string, timeRes string, dateTimeRes string) {
	if addDate {
		if d.currentDate == "" {
			d.currentDate = d.timeNow().Format("02/01/2006")
			go d.updateCurrentDateEveryDay()
		}

		dateRes = d.currentDate
	}

	if addTime {
		if d.currentTime == "" {
			d.currentTime = d.timeNow().Format("15:04:05")
			go d.updateCurrentTimeEverySecond()
		}

		timeRes = d.currentTime
	}

	if addDate && addTime {
		return "", "", dateRes + " " + timeRes
	}

	return dateRes, timeRes, ""
}

// updateCurrentDateEveryDay synchronizes with the system clock and updates the DateTimePrinter's
// currentDate property every full Day.
func (d *DateTimePrinter) updateCurrentDateEveryDay() {
	now := d.timeNow()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	initialDelay := now.Sub(midnight) // Time since midnight
	time.Sleep(initialDelay)

	daysTicker := time.NewTicker(1 * (time.Hour * 24))

	for t := range daysTicker.C {
		d.currentDate = t.Format("02/01/2006")
	}
}

// updateCurrentTimeEverySecond synchronizes with the system clock and updates the DateTimePrinter's
// currentTime property every full second.
func (d *DateTimePrinter) updateCurrentTimeEverySecond() {
	initialDelay := 1*time.Second - time.Duration(d.timeNow().Nanosecond())*time.Nanosecond
	time.Sleep(initialDelay)

	secondsTicker := time.NewTicker(1 * time.Second)

	for t := range secondsTicker.C {
		d.currentTime = t.Format("15:04:05")
	}
}

// NewDateTimePrinter initializes DateTimePrinter with default timeNow function.
func NewDateTimePrinter() *DateTimePrinter {
	return &DateTimePrinter{timeNow: time.Now}
}
