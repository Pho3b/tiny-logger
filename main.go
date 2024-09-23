package main

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/logs"
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/configs"
	ll "gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

func main() {
	logger := logs.NewLogger().
		SetConfigs(
			&configs.TLConfigs{
				AddDateTime:  true,
				EnableColors: false,
				Parser:       "",
				LogLvl:       ll.LogLevel{},
			},
		).
		SetLogLvl(ll.DebugLvlName)

	logger.Debug("This is my Debug Log", "Esitating")
}
