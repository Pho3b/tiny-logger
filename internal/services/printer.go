package services

import (
	"bytes"
	"os"

	c "github.com/Pho3b/tiny-logger/logs/colors"
	"github.com/Pho3b/tiny-logger/logs/log_level"
	s "github.com/Pho3b/tiny-logger/shared"
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
func (p *Printer) RetrieveColorsFromLogLevel(enableColors bool, logLevelName log_level.LogLvlName) [2]c.Color {
	var res = [2]c.Color{"", ""}

	if enableColors {
		switch logLevelName {
		case log_level.FatalErrorLvlName:
			res[0] = c.Magenta
		case log_level.ErrorLvlName:
			res[0] = c.Red
		case log_level.WarnLvlName:
			res[0] = c.Yellow
		case log_level.InfoLvlName:
			res[0] = c.Cyan
		case log_level.DebugLvlName:
			res[0] = c.Gray
		}

		res[1] = c.Reset
	}

	return res
}

func NewPrinter() Printer {
	return Printer{}
}
