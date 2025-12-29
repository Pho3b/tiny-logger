package services

import (
	"bytes"
	"context"
	"os"
	"sync"
	"time"

	"github.com/pho3b/tiny-logger/shared"
)

type LogsBuffer struct {
	StopLogs            context.CancelFunc
	printer             Printer
	msgBuf              *bytes.Buffer
	mtx                 sync.Mutex
	ctx                 context.Context
	logFile             *os.File
	bufferFlushInterval time.Duration
	outType             shared.OutputType
}

func (l *LogsBuffer) AddLog(buf *bytes.Buffer) {
	l.mtx.Lock()
	l.msgBuf.Write(buf.Bytes())
	l.mtx.Unlock()
}

func (l *LogsBuffer) UpdateLogFile(file *os.File) {
	l.logFile = file

	if file == nil {
		l.outType = shared.StdOutput
	} else {
		l.outType = shared.FileOutput
	}
}

func (l *LogsBuffer) GetFlushInterval() time.Duration {
	return l.bufferFlushInterval
}

// SetBufferFlushInterval sets the interval at which the logs buffer will flush its logs to the output file.
// If the given interval is <= 0, the buffered log is stopped and logs will be printed in real time.
func (l *LogsBuffer) SetBufferFlushInterval(interval time.Duration) {
	if interval <= 0 {
		l.bufferFlushInterval = 0
		l.StopLogs()

		return
	}

	// context is not nil OR DONE
	if l.ctx == nil || l.ctx.Err() != nil {
		ctx, cancel := context.WithCancel(context.Background())
		l.ctx = ctx
		l.StopLogs = cancel
	}

	l.bufferFlushInterval = interval
	l.startInternalTicker()
}

func (l *LogsBuffer) startInternalTicker() {
	ticker := time.NewTicker(l.bufferFlushInterval)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-l.ctx.Done():
				l.FlushLogs()
				_, _ = os.Stdout.Write([]byte("Stopping logs buffer..."))
				return

			case <-ticker.C:
				l.FlushLogs()
			}
		}
	}()
}

func (l *LogsBuffer) FlushLogs() {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if l.msgBuf.Len() == 0 {
		return
	}

	l.printer.PrintLog(l.outType, l.msgBuf, l.logFile)
	l.msgBuf.Reset()
}

func NewLogsBuffer(flushInterval time.Duration, logFile *os.File, printer Printer) *LogsBuffer {
	outType := shared.StdOutput
	if logFile != nil {
		outType = shared.FileOutput
	}

	b := &LogsBuffer{
		msgBuf:              &bytes.Buffer{},
		logFile:             logFile,
		bufferFlushInterval: flushInterval,
		printer:             printer,
		outType:             outType,
	}

	b.msgBuf.Grow(10000000)
	return b
}
