package services

import (
	"bytes"
	"context"
	"os"
	"sync"
	"time"
)

type LogsBuffer struct {
	StopLogs     context.CancelFunc
	msgBuf       *bytes.Buffer
	mtx          sync.Mutex
	logsInterval time.Duration
	ctx          context.Context
}

func (l *LogsBuffer) AddLogFrom(buf *bytes.Buffer) {
	l.mtx.Lock()
	l.msgBuf.Write(buf.Bytes())
	l.mtx.Unlock()
}

func (l *LogsBuffer) startInternalTicker() {
	ticker := time.NewTicker(l.logsInterval)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-l.ctx.Done():
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

	// TODO: Update it to actually follow the Logger configurations
	_, _ = os.Stdout.Write(l.msgBuf.Bytes())
	l.msgBuf.Reset()
}

func NewLogsBuffer() *LogsBuffer {
	ctx, cancel := context.WithCancel(context.Background())

	b := &LogsBuffer{
		StopLogs:     cancel,
		ctx:          ctx,
		msgBuf:       &bytes.Buffer{},
		logsInterval: (time.Second * 2),
	}

	b.msgBuf.Grow(10000000)
	b.startInternalTicker()

	return b
}
