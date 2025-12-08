package services

import (
	"bytes"
	"os"

	s "github.com/pho3b/tiny-logger/shared"
)

// PrintLog prints the given msgBuffer to the given outputType (stdout or stderr).
// If 'file' is not nil, the message is written to the file.
func PrintLog(outType s.OutputType, msgBuffer *bytes.Buffer, file *os.File) {
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
