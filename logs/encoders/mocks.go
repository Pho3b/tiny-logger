package encoders

import (
	"gitlab.com/docebo/libraries/go/tiny-logger/logs/log_level"
)

type LoggerConfigMock struct {
	DateEnabled   bool
	TimeEnabled   bool
	ColorsEnabled bool
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
