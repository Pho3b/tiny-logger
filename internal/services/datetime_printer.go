package services

import (
	"strconv"
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
	timeNow       func() time.Time // Function to get current time, allows mocking for tests
	currentFormat atomic.Value
	currentDate   atomic.Value
	currentTime   atomic.Value
	currentUnix   atomic.Value
	dateOnce      sync.Once
	timeOnce      sync.Once
	unixOnce      sync.Once
}

// RetrieveDateTime returns the current date, time, and combined/unix string based on the configuration.
// Returns: (dateString, timeString, combinedOrUnixString).
// If the format is UnixTimestamp, the timestamp is returned as the third value, ignoring the boolean flags.
// Otherwise, 'addDate' and 'addTime' control which components are generated.
func (d *DateTimePrinter) RetrieveDateTime(addDate, addTime bool) (string, string, string) {
	var dateRes, timeRes string
	currentFmt := d.currentFormat.Load().(s.DateTimeFormat)

	if currentFmt == s.UnixTimestamp && (addDate || addTime) {
		d.unixOnce.Do(func() {
			d.currentUnix.Store(strconv.FormatInt(d.timeNow().Unix(), 10))
			go d.updateCurrentUnixEverySecond()
		})

		return "", "", d.currentUnix.Load().(string)
	}

	if addDate {
		d.dateOnce.Do(func() {
			d.currentDate.Store(d.timeNow().Format(dateFormat[currentFmt]))
			go d.updateCurrentDateEveryDay()
		})

		dateRes = d.currentDate.Load().(string)
	}

	if addTime {
		d.timeOnce.Do(func() {
			d.currentTime.Store(d.timeNow().Format(timeFormat[currentFmt]))
			go d.updateCurrentTimeEverySecond()
		})

		timeRes = d.currentTime.Load().(string)
	}

	if addDate && addTime {
		return "", "", dateRes + " " + timeRes
	}

	return dateRes, timeRes, ""
}

// UpdateDateTimeFormat updates the DateTimePrinter's currentFormat property and updates the currentDate and
// currentTime properties accordingly.
func (d *DateTimePrinter) UpdateDateTimeFormat(format s.DateTimeFormat) {
	now := d.timeNow()

	d.currentFormat.Store(format)
	d.currentDate.Store(now.Format(dateFormat[format]))
	d.currentTime.Store(now.Format(timeFormat[format]))
}

// updateCurrentDateEveryDay synchronizes with the system clock and updates the DateTimePrinter's
// currentDate property every midnight.
func (d *DateTimePrinter) updateCurrentDateEveryDay() {
	for {
		now := d.timeNow()
		currentFmt := d.currentFormat.Load().(s.DateTimeFormat)
		d.currentDate.Store(now.Format(dateFormat[currentFmt]))

		// computing next midnight in local time zone
		nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		time.Sleep(time.Until(nextMidnight))
	}
}

// updateCurrentTimeEverySecond synchronizes with the system clock and updates the DateTimePrinter's
// currentTime property every full second.
func (d *DateTimePrinter) updateCurrentTimeEverySecond() {
	for {
		now := d.timeNow()
		currentFmt := d.currentFormat.Load().(s.DateTimeFormat)
		d.currentTime.Store(now.Format(timeFormat[currentFmt]))

		nextSecond := now.Truncate(time.Second).Add(time.Second)
		time.Sleep(time.Until(nextSecond))
	}
}

// updateCurrentUnixEverySecond synchronizes with the system clock and updates the DateTimePrinter's
// currentTime property every full second.
func (d *DateTimePrinter) updateCurrentUnixEverySecond() {
	for {
		now := d.timeNow()
		d.currentUnix.Store(strconv.FormatInt(now.Unix(), 10))

		nextSecond := now.Truncate(time.Second).Add(time.Second)
		time.Sleep(time.Until(nextSecond))
	}
}

// NewDateTimePrinter initializes DateTimePrinter with the default timeNow function.
func NewDateTimePrinter() *DateTimePrinter {
	dateTimePrinter := &DateTimePrinter{timeNow: time.Now}
	dateTimePrinter.currentFormat.Store(s.IT)

	return dateTimePrinter
}
