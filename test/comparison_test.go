package test

import (
	"io"
	"os"
	"testing"

	"github.com/pho3b/tiny-logger/logs"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// devNullFile is used for tiny-logger which requires an *os.File
var devNullFile *os.File

func init() {
	var err error
	// Open /dev/null (or NUL on Windows) to discard output for tiny-logger
	devNullFile, err = os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		panic(err)
	}
}

// 1. Tiny Logger (Project)
func BenchmarkTinyLogger(b *testing.B) {
	// Initialize tiny-logger
	logger := logs.NewLogger()
	// Redirect to /dev/null to measure logger overhead only, not I/O
	logger.SetLogFile(devNullFile)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Simulating a log with a message and a few fields
		logger.Info("Benchmark message", "iteration", i, "active", true)
	}
}

// 2. Logrus (Standard-like popular logger)
func BenchmarkLogrus(b *testing.B) {
	logger := logrus.New()
	logger.Out = io.Discard
	logger.SetFormatter(&logrus.TextFormatter{DisableColors: true})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"iteration": i,
			"active":    true,
		}).Info("Benchmark message")
	}
}

// 3. Zerolog (Zero Allocation Logger)
func BenchmarkZerolog(b *testing.B) {
	logger := zerolog.New(io.Discard)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info().
			Any("iteration", i).
			Any("active", true).
			Any("Benchmark message", "")
	}
}

// 4. Zap (Uber's fast logger)
func BenchmarkZap(b *testing.B) {
	// Configure Zap to discard output
	encoderConfig := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(io.Discard),
		zap.InfoLevel,
	)
	logger := zap.New(core)
	defer logger.Sync()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("Benchmark message",
			zap.Any("iteration", i),
			zap.Any("active", true),
		)
	}
}
