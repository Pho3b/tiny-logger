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

var (
	dateTimePrinterInstance *DateTimePrinter
	dateTimePrinterOnce     sync.Once
)

type DateTimePrinter struct {
	timeNow     func() time.Time
	cachedDates [3]atomic.Value
	cachedTimes [3]atomic.Value
	currentUnix atomic.Value
}

// RetrieveDateTime now accepts the desired format
func (d *DateTimePrinter) RetrieveDateTime(fmt s.DateTimeFormat, addDate, addTime bool) (string, string, string) {
	if fmt == s.UnixTimestamp {
		return "", "", d.currentUnix.Load().(string)
	}

	var dateRes, timeRes string

	if addDate {
		dateRes = d.cachedDates[fmt].Load().(string)
	}

	if addTime {
		timeRes = d.cachedTimes[fmt].Load().(string)
	}

	return dateRes, timeRes, ""
}

// init initializes the current timestamp and cached formatted strings,
// then starts background goroutines to keep them updated.
func (d *DateTimePrinter) init() {
	now := d.timeNow()
	d.currentUnix.Store(strconv.FormatInt(now.Unix(), 10))

	for i := 0; i < 3; i++ {
		fmt := s.DateTimeFormat(i)
		d.cachedDates[i].Store(now.Format(dateFormat[fmt]))
		d.cachedTimes[i].Store(now.Format(timeFormat[fmt]))
	}

	go d.loopUpdateDate()
	go d.loopUpdateTime()
}

// loopUpdateTime updates all time formats, the unix timestamp every second,
// and refreshes the date format if the day has changed.
func (d *DateTimePrinter) loopUpdateTime() {
	// Initialize lastDay with a value that forces an update on the first iteration
	lastDay := -1

	for {
		now := d.timeNow()

		// 1. Update Time Formats (Always)
		for i := 0; i < 3; i++ {
			fmt := s.DateTimeFormat(i)
			d.cachedTimes[i].Store(now.Format(timeFormat[fmt]))
		}

		// 2. Update Unix Timestamp (Always)
		d.currentUnix.Store(strconv.FormatInt(now.Unix(), 10))

		// 3. Update Date Formats (Only if day changed)
		if currentDay := now.Day(); currentDay != lastDay {
			for i := 0; i < 3; i++ {
				fmt := s.DateTimeFormat(i)
				d.cachedDates[i].Store(now.Format(dateFormat[fmt]))
			}

			lastDay = currentDay
		}

		nextSecond := now.Truncate(time.Second).Add(time.Second)
		time.Sleep(time.Until(nextSecond))
	}
}

// loopUpdateDate updates all date formats every 10 mins
func (d *DateTimePrinter) loopUpdateDate() {
	for {
		now := d.timeNow()

		for i := 0; i < 3; i++ {
			fmt := s.DateTimeFormat(i)
			d.cachedDates[i].Store(now.Format(dateFormat[fmt]))
		}

		time.Sleep(time.Minute * 10)
	}
}

// GetDateTimePrinter returns the singleton instance.
func GetDateTimePrinter() *DateTimePrinter {
	dateTimePrinterOnce.Do(
		func() {
			dateTimePrinterInstance = &DateTimePrinter{timeNow: time.Now}
			dateTimePrinterInstance.init()
		},
	)

	return dateTimePrinterInstance
}
