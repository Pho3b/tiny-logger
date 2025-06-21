package test

import (
	"fmt"
	"github.com/pho3b/tiny-logger/logs"
	"github.com/pho3b/tiny-logger/shared"
	"testing"
)

func BenchmarkDefaultEncoderAllPropertiesDisabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.DefaultEncoderType).ShowLogLevel(false)

	for i := 0; i < b.N; i++ {
		logger.Debug("DEFAULT encoder", "all-properties-enabled", false, "id", 2)
	}

	fmt.Print("BenchmarkDefaultEncoderAllPropertiesDisabled:")
}

func BenchmarkDefaultEncoderAllPropertiesEnabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.DefaultEncoderType).
		ShowLogLevel(true).
		AddDateTime(true).
		EnableColors(true)

	for i := 0; i < b.N; i++ {
		logger.Debug("DEFAULT encoder", "all-properties-enabled", true, "id", 2)
	}

	fmt.Print("BenchmarkDefaultEncoderAllPropertiesEnabled:")
}

func BenchmarkJsonEncoderAllPropertiesDisables(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.JsonEncoderType).ShowLogLevel(false)

	for i := 0; i < b.N; i++ {
		logger.Debug("JSON encoder", "all-properties-enabled", false, "id", 2)
	}

	fmt.Print("BenchmarkJsonEncoderAllPropertiesDisables:")
}

func BenchmarkJsonEncoderAllPropertiesEnabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.JsonEncoderType).
		ShowLogLevel(true).
		AddDateTime(true)

	for i := 0; i < b.N; i++ {
		logger.Debug("JSON encoder", "all-properties-enabled", true, "id", 2)
	}

	fmt.Print("BenchmarkJsonEncoderAllPropertiesEnabled:")
}

func BenchmarkYamlEncoderAllPropertiesDisables(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.YamlEncoderType).ShowLogLevel(false)

	for i := 0; i < b.N; i++ {
		logger.Debug("YAML encoder", "all-properties-enabled", false, "id", 2)
	}

	fmt.Print("BenchmarkYamlEncoderAllPropertiesDisables:")
}

func BenchmarkYamlEncoderAllPropertiesEnabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.YamlEncoderType).
		ShowLogLevel(true).
		AddDateTime(true)

	for i := 0; i < b.N; i++ {
		logger.Debug("YAML encoder", "all-properties-enabled", true, "id", 2)
	}

	fmt.Print("BenchmarkYamlEncoderAllPropertiesEnabled:")
}
