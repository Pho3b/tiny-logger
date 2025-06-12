package services

import (
	"sync"
	"sync/atomic"
	"time"
)

type DateTimePrinter struct {
	timeNow     func() time.Time // Function to get current time, allows mocking for tests
	currentDate atomic.Value
	currentTime atomic.Value
	dateOnce    sync.Once
	timeOnce    sync.Once
}

// RetrieveDateTime returns formatted date and/or time strings based on input flags.
// - If addDate is true, the method returns the current date in "DD/MM/YYYY" format.
// - If addTime is true, it returns the current time in "HH:MM:SS" format.
// - If both addDate and addTime are true, dateTimeRes is returned as a unified string.
// - If neither addDate nor addTime is true, empty strings are returned.
func (d *DateTimePrinter) RetrieveDateTime(addDate, addTime bool) (string, string, string) {
	var dateRes, timeRes string

	if addDate {
		cDate := d.currentDate.Load()

		if cDate == nil {
			d.currentDate.Store(d.timeNow().Format("02/01/2006"))
			d.dateOnce.Do(func() {
				go d.updateCurrentDateEveryDay()
			})

			cDate = d.currentDate.Load()
		}

		dateRes = cDate.(string)
	}

	if addTime {
		cTime := d.currentTime.Load()

		if cTime == nil {
			d.currentTime.Store(d.timeNow().Format("15:04:05"))
			d.timeOnce.Do(func() {
				go d.updateCurrentTimeEverySecond()
			})

			cTime = d.currentTime.Load()
		}

		timeRes = cTime.(string)
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

	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for t := range ticker.C {
		d.currentDate.Store(t.Format("02/01/2006"))
	}
}

// updateCurrentTimeEverySecond synchronizes with the system clock and updates the DateTimePrinter's
// currentTime property every full second.
func (d *DateTimePrinter) updateCurrentTimeEverySecond() {
	initialDelay := 1*time.Second - time.Duration(d.timeNow().Nanosecond())*time.Nanosecond
	time.Sleep(initialDelay)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		d.currentTime.Store(t.Format("15:04:05"))
	}
}

// NewDateTimePrinter initializes DateTimePrinter with default timeNow function.
func NewDateTimePrinter() DateTimePrinter {
	return DateTimePrinter{timeNow: time.Now}
}
