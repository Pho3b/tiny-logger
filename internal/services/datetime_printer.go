package services

import (
	"fmt"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/configs"
	"time"
)

type DateTimePrinter struct {
}

func (d *DateTimePrinter) PrintDateTime(conf *configs.TLConfigs) string {
	dateTime := ""

	if conf.AddDateTime {
		dateTime = fmt.Sprintf("[%s]", time.Now().Format("02/01/2006 15:04:05"))
	}

	return dateTime
}
