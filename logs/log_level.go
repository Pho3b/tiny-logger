package logs

type LogLevel struct {
	lvl         int8
	envVariable string
}

func (l *LogLevel) LvlName() string {
	return logLvlIntToName[l.lvl]
}

func (l *LogLevel) LvlIntValue() int8 {
	return l.lvl
}
