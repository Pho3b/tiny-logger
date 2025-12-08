package test

import (
	"os"
	"time"

	"github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/pho3b/tiny-logger/shared"
)

type LoggerConfigMock struct {
	DateEnabled   bool
	TimeEnabled   bool
	ColorsEnabled bool
	ShowLogLevel  bool
}

func (m *LoggerConfigMock) GetLogLvlName() log_level.LogLvlName {
	return log_level.DebugLvlName
}

func (m *LoggerConfigMock) GetLogLvlIntValue() int8 {
	return log_level.DebugLvl
}

func (m *LoggerConfigMock) GetDateTimeEnabled() (dateEnabled bool, timeEnabled bool) {
	return m.DateEnabled, m.TimeEnabled
}
func (m *LoggerConfigMock) GetColorsEnabled() bool {
	return m.ColorsEnabled
}

func (m *LoggerConfigMock) GetEncoderType() shared.EncoderType {
	return shared.DefaultEncoderType
}

func (m *LoggerConfigMock) GetShowLogLevel() bool {
	return m.ShowLogLevel
}

func (m *LoggerConfigMock) GetLogFile() *os.File {
	return nil
}

func (m *LoggerConfigMock) GetDateTimeFormat() shared.DateTimeFormat {
	return shared.IT
}

func (m *LoggerConfigMock) GetBufferFlushInterval() time.Duration {
	return 1 * time.Second
}
