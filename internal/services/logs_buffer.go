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
	StopLogs context.CancelFunc
	printer  PrinterService
	msgBuf   *bytes.Buffer
	mtx      sync.Mutex
	configs  shared.LoggerConfigsInterface
	ctx      context.Context
	logFile  *os.File
}

func (l *LogsBuffer) AddLog(buf *bytes.Buffer) {
	l.mtx.Lock()
	l.msgBuf.Write(buf.Bytes())
	l.mtx.Unlock()
}

func (l *LogsBuffer) UpdateLogFile(file *os.File) {
	l.logFile = file
}

func (l *LogsBuffer) startInternalTicker() {
	ticker := time.NewTicker(l.configs.GetBufferFlushInterval())

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-l.ctx.Done():
				l.flushLogs()
				_, _ = os.Stdout.Write([]byte("Stopping logs buffer..."))
				return

			case <-ticker.C:
				l.flushLogs()
			}
		}
	}()
}

func (l *LogsBuffer) flushLogs() {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if l.msgBuf.Len() == 0 {
		return
	}

	l.printer.PrintLog(shared.StdOutput, l.msgBuf, l.logFile)
	l.msgBuf.Reset()
}

func NewLogsBuffer(loggerConfigs shared.LoggerConfigsInterface) *LogsBuffer {
	ctx, cancel := context.WithCancel(context.Background())

	b := &LogsBuffer{
		StopLogs: cancel,
		ctx:      ctx,
		msgBuf:   &bytes.Buffer{},
		configs:  loggerConfigs,
		logFile:  loggerConfigs.GetLogFile(),
	}

	b.msgBuf.Grow(10000000)
	b.startInternalTicker()

	return b
}
