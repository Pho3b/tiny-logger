package services

import (
	"sync"
	"sync/atomic"
	"time"

	s "github.com/pho3b/tiny-logger/shared"
)

var dateFormat = map[s.DateTimeFormat]string{
	s.IT: "02/01/2006",
	s.US: "01/02/2006",
	s.JP: "2006/01/02",
}

var timeFormat = map[s.DateTimeFormat]string{
	s.IT: "15:04:05",
	s.US: "03:04:05 PM",
	s.JP: "03:04:05 PM",
}

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
func (d *DateTimePrinter) RetrieveDateTime(addDate, addTime bool, format s.DateTimeFormat) (string, string, string) {
	var dateRes, timeRes string
	now := d.timeNow()

	if addDate {
		cDate := d.currentDate.Load()

		if cDate == nil {
			d.currentDate.Store(now.Format(dateFormat[format]))
			d.dateOnce.Do(func() {
				go d.updateCurrentDateEveryDay(format)
			})

			cDate = d.currentDate.Load()
		}

		dateRes = cDate.(string)
	}

	if addTime {
		cTime := d.currentTime.Load()

		if cTime == nil {
			d.currentTime.Store(now.Format(timeFormat[format]))
			d.timeOnce.Do(func() {
				go d.updateCurrentTimeEverySecond(format)
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
// currentDate property every midnight.
func (d *DateTimePrinter) updateCurrentDateEveryDay(format s.DateTimeFormat) {
	for {
		now := d.timeNow()
		d.currentDate.Store(now.Format(dateFormat[format]))

		// computing next midnight in local time zone
		nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		time.Sleep(time.Until(nextMidnight))
	}
}

// updateCurrentTimeEverySecond synchronizes with the system clock and updates the DateTimePrinter's
// currentTime property every full second.
func (d *DateTimePrinter) updateCurrentTimeEverySecond(format s.DateTimeFormat) {
	for {
		now := d.timeNow()
		d.currentTime.Store(now.Format(timeFormat[format]))

		nextSecond := now.Truncate(time.Second).Add(time.Second)
		time.Sleep(time.Until(nextSecond))
	}
}

// NewDateTimePrinter initializes DateTimePrinter with the default timeNow function.
func NewDateTimePrinter() DateTimePrinter {
	return DateTimePrinter{timeNow: time.Now}
}
