package test

import (
	"testing"

	"github.com/Pho3b/tiny-logger/logs"
	"github.com/Pho3b/tiny-logger/shared"
)

func BenchmarkDefaultEncoderAllPropertiesDisabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.DefaultEncoderType).ShowLogLevel(false).SetLogFile(initDevNullFile())

	for i := 0; i < b.N; i++ {
		logger.Debug("DEFAULT encoder", "all-properties-enabled", false, "id", i)
	}
}

func BenchmarkDefaultEncoderAllPropertiesEnabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.DefaultEncoderType).
		ShowLogLevel(true).
		AddDateTime(true).
		EnableColors(true).
		SetLogFile(initDevNullFile())

	for i := 0; i < b.N; i++ {
		logger.Debug("DEFAULT encoder", "all-properties-enabled", true, "id", i)
	}
}

func BenchmarkJsonEncoderAllPropertiesDisabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.JsonEncoderType).ShowLogLevel(false).SetLogFile(initDevNullFile())

	for i := 0; i < b.N; i++ {
		logger.Debug("JSON encoder", "all-properties-enabled", false, "id", i)
	}
}

func BenchmarkJsonEncoderAllPropertiesEnabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.JsonEncoderType).
		ShowLogLevel(true).
		AddDateTime(true).
		SetLogFile(initDevNullFile())

	for i := 0; i < b.N; i++ {
		logger.Debug("JSON encoder", "all-properties-enabled", true, "id", i)
	}
}

func BenchmarkYamlEncoderAllPropertiesDisabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.YamlEncoderType).ShowLogLevel(false).SetLogFile(initDevNullFile())

	for i := 0; i < b.N; i++ {
		logger.Debug("YAML encoder", "all-properties-enabled", false, "id", i)
	}
}

func BenchmarkYamlEncoderAllPropertiesEnabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.YamlEncoderType).
		ShowLogLevel(true).
		AddDateTime(true).
		SetLogFile(initDevNullFile())

	for i := 0; i < b.N; i++ {
		logger.Debug("YAML encoder", "all-properties-enabled", true, "id", i)
	}
}
