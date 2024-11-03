package main

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/logs"
	ll "gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
	"gitlab.com/docebo/libraries/go/tiny-logger/shared"
)

func main() {
	logger := logs.NewLogger().
		SetLogLvl(ll.ErrorLvlName).
		EnableColors(true).
		AddTime(true).AddDate(true)

	logger.Error("This is my Error log", "Test second arg")
	logger.SetLogLvl(ll.DebugLvlName)
	logger.Warn("This is my Warn log", "Test second arg")
	logger.Debug("This is my Debug log", "Test second arg")

	logger.SetEncoder(shared.JsonEncoderType)
	logger.Error("This is my Error log", "Test second arg")
	logger.SetLogLvl(ll.DebugLvlName)
	logger.Warn("This is my Warn log", "Test second arg")
	logger.Debug("This is my Debug log", "Test second arg")

	logger.SetEncoder(shared.YamlEncoderType)
	logger.Error("This is my Error log", "Test second arg")
	logger.SetLogLvl(ll.DebugLvlName)
	logger.Warn("This is my Warn log", "Test second arg")
	logger.Debug("This is my Debug log", "Test second arg")
}
