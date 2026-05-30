// Package comparison hosts the benchmarks that compare tiny-logger against third-party loggers.
// It lives in its own Go module so that zerolog/zap/logrus (and their transitive requirements)
// never become dependencies of the main tiny-logger module, keeping the library at its real
// minimum Go version and dependency-free.
package comparison

import (
	"os"
	"testing"

	"github.com/Pho3b/tiny-logger/logs"
	"github.com/Pho3b/tiny-logger/shared"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// initDevNullFile opens os.DevNull to discard output, isolating logger overhead from real I/O.
func initDevNullFile() *os.File {
	devNullFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		panic(err)
	}

	return devNullFile
}

// 1. Tiny Logger (Project)
func BenchmarkTinyLogger(b *testing.B) {
	// Initialize tiny-logger
	logger := logs.NewLogger().SetEncoder(shared.JsonEncoderType).AddDateTime(true)
	// Redirect to /dev/null to measure logger overhead only, not I/O
	logger.SetLogFile(initDevNullFile())

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
	logger.Out = initDevNullFile()
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
	logger := zerolog.New(initDevNullFile())

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info().
			Any("iteration", i).
			Any("active", true).
			Msg("Benchmark message")
	}
}

// 4. Zap (Uber's fast logger)
func BenchmarkZap(b *testing.B) {
	// Configure Zap to discard output
	encoderConfig := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(initDevNullFile()),
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
