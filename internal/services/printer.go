package services

import (
	"bytes"
	"os"

	c "github.com/pho3b/tiny-logger/logs/colors"
	"github.com/pho3b/tiny-logger/logs/log_level"
	s "github.com/pho3b/tiny-logger/shared"
)

type Printer struct {
}

// PrintLog prints the given msgBuffer to the given outputType (stdout or stderr).
// If 'file' is not nil AND (outType == FileOutput), the message is written to the file.
func (p *Printer) PrintLog(outType s.OutputType, msgBuffer *bytes.Buffer, file *os.File) {
	var err error

	switch outType {
	case s.StdOutput:
		_, err = os.Stdout.Write(msgBuffer.Bytes())
	case s.StdErrOutput:
		_, err = os.Stderr.Write(msgBuffer.Bytes())
	case s.FileOutput:
		if file == nil {
			_, _ = os.Stderr.Write([]byte("tiny-logger-err: given out file is nil"))
			return
		}

		_, err = file.Write(msgBuffer.Bytes())
	}

	if err != nil {
		_, _ = os.Stderr.Write([]byte("tiny-logger-err: " + err.Error() + "\n"))
	}
}

// RetrieveColorsFromLogLevel returns an array of colors as strings to be used in log output based on a given log level.
// if enableColors is false, it returns an array of empty strings.
func (p *Printer) RetrieveColorsFromLogLevel(enableColors bool, logLevelInt int8) []c.Color {
	var res = []c.Color{"", ""}

	if enableColors {
		switch logLevelInt {
		case log_level.FatalErrorLvl:
			res[0] = c.Magenta
		case log_level.ErrorLvl:
			res[0] = c.Red
		case log_level.WarnLvl:
			res[0] = c.Yellow
		case log_level.InfoLvl:
			res[0] = c.Cyan
		case log_level.DebugLvl:
			res[0] = c.Gray
		}

		res[1] = c.Reset
	}

	return res
}

func NewPrinter() Printer {
	return Printer{}
}
