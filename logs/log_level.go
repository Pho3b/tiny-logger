package logs

// LogLvlName is the Enum representing the possible Log Levels.
type LogLvlName string

const (
	ErrorLvlName LogLvlName = "ERROR"
	WarnLvlName  LogLvlName = "WARN"
	InfoLvlName  LogLvlName = "INFO"
	DebugLvlName LogLvlName = "DEBUG"
)

// Log Level INT8 Constants
const (
	ErrorLvl = int8(0)
	WarnLvl  = int8(1)
	InfoLvl  = int8(2)
	DebugLvl = int8(3)
)

// Log Level INT to STRING map
var logLvlIntToName = map[int8]LogLvlName{
	ErrorLvl: ErrorLvlName,
	WarnLvl:  WarnLvlName,
	InfoLvl:  InfoLvlName,
	DebugLvl: DebugLvlName,
}

// Log Level STRING to INT map
var logLvlNameToInt = map[LogLvlName]int8{
	ErrorLvlName: ErrorLvl,
	WarnLvlName:  WarnLvl,
	InfoLvlName:  InfoLvl,
	DebugLvlName: DebugLvl,
}

type LogLevel struct {
	lvl         int8
	envVariable string
}

func (l *LogLevel) LvlName() LogLvlName {
	return logLvlIntToName[l.lvl]
}

func (l *LogLevel) LvlIntValue() int8 {
	return l.lvl
}
