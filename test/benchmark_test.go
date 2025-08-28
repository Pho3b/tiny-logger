package test

import (
	"fmt"
	"testing"

	"github.com/pho3b/tiny-logger/logs"
	"github.com/pho3b/tiny-logger/shared"
)

func BenchmarkDefaultEncoderAllPropertiesDisabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.DefaultEncoderType).ShowLogLevel(false)

	for i := 0; i < b.N; i++ {
		logger.Debug("DEFAULT encoder", "all-properties-enabled", false, "id", 2)
	}

	fmt.Print("Default_Encoder_All_Properties_Disabled:")
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

	fmt.Print("Default_Encoder_All_Properties_Enabled: ")
}

func BenchmarkJsonEncoderAllPropertiesDisabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.JsonEncoderType).ShowLogLevel(false)

	for i := 0; i < b.N; i++ {
		logger.Debug("JSON encoder", "all-properties-enabled", false, "id", 2)
	}

	fmt.Print("Json_Encoder_All_Properties_Disabled: ")
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

	fmt.Print("Json_Encoder_All_Properties_Enabled: ")
}

func BenchmarkYamlEncoderAllPropertiesDisabled(b *testing.B) {
	b.ReportAllocs()

	logger := logs.NewLogger().
		SetEncoder(shared.YamlEncoderType).ShowLogLevel(false)

	for i := 0; i < b.N; i++ {
		logger.Debug("YAML encoder", "all-properties-enabled", false, "id", 2)
	}

	fmt.Print("Yaml_Encoder_All_Properties_Disabled: ")
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

	fmt.Print("Yaml_Encoder_All_Properties_Enabled: ")
}
