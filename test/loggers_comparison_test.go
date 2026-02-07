package test

import (
	"testing"

	"github.com/Pho3b/tiny-logger/logs"
	"github.com/Pho3b/tiny-logger/shared"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
