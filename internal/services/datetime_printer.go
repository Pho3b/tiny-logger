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
	timeNow func() time.Time
	// Arrays indexed by s.DateTimeFormat (int8)
	// Assuming 0=IT, 1=JP, 2=US. UnixTimestamp is handled separately.
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
	go d.loopUpdateUnix()
}

// loopUpdateTime updates all time formats every second
func (d *DateTimePrinter) loopUpdateTime() {
	for {
		now := d.timeNow()

		for i := 0; i < 3; i++ {
			fmt := s.DateTimeFormat(i)
			d.cachedTimes[i].Store(now.Format(timeFormat[fmt]))
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

// updateCurrentTime synchronizes with the system clock and updates the DateTimePrinter's
// currentTime property every full second.
func (d *DateTimePrinter) loopUpdateUnix() {
	for {
		now := d.timeNow()
		d.currentUnix.Store(strconv.FormatInt(now.Unix(), 10))

		nextSecond := now.Truncate(time.Second).Add(time.Second)
		time.Sleep(time.Until(nextSecond))
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
