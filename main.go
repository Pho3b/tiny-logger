package main

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/logs"
	ll "gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

func main() {
	logger := logs.NewLogger().
		SetLogLvl(ll.DebugLvlName).
		SetAddDateTime(true).
		SetEnableColors(false)

	logger.Debug("This is my Debug Log", "Esitating")
}
