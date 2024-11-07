package main

import "github.com/pho3b/tiny-logger/logs"

func main() {
	logger := logs.NewLogger().EnableColors(true).AddDateTime(true)

	logger.Error("test", nil)
	logger.Warn("test", nil)

	logger.AddDateTime(false)
	logger.Error("test", nil)
	logger.Warn("test", nil)

	logger.AddTime(true)
	logger.Error("test", nil)
	logger.Warn("test", nil)
}
