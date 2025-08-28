package services

import (
	"bytes"
	"testing"
)

func TestYamlMarshaler_Marshal(t *testing.T) {
	tests := []struct {
		name     string
		entry    YamlLogEntry
		expected string
	}{
		{
			name: "basic message only",
			entry: YamlLogEntry{
				Message: "test message",
			},
			expected: "msg: test message\n",
		},
		{
			name: "message with level",
			entry: YamlLogEntry{
				Level:   "INFO",
				Message: "test message",
			},
			expected: "level: INFO\nmsg: test message\n",
		},
		{
			name: "full log entry",
			entry: YamlLogEntry{
				Level:    "DEBUG",
				Date:     "2024-03-21",
				Time:     "15:04:05",
				DateTime: "2024-03-21T15:04:05",
				Message:  "full test message",
			},
			expected: "level: DEBUG\ndate: 2024-03-21\ntime: 15:04:05\ndatetime: 2024-03-21T15:04:05\nmsg: full test message\n",
		},
		{
			name: "with simple extras",
			entry: YamlLogEntry{
				Level:   "INFO",
				Message: "test with extras",
				Extras:  []any{"key1", "value1", "key2", 42, "int64", int64(45)},
			},
			expected: "level: INFO\nmsg: test with extras\nextras:\n  key1: value1\n  key2: 42\n  int64: 45\n",
		},
		{
			name: "with special character extras",
			entry: YamlLogEntry{
				Message: "test with special chars",
				Extras:  []any{"key:1", "value: with colon", "key-2", "value with spaces"},
			},
			expected: "msg: test with special chars\nextras:\n  \"key:1\": \"value: with colon\"\n  \"key-2\": \"value with spaces\"\n",
		},
		{
			name: "with odd number of extras",
			entry: YamlLogEntry{
				Message: "odd extras",
				Extras:  []any{"key1", "value1", "key2"},
			},
			expected: "msg: odd extras\nextras:\n  key1: value1\n  key2: null\n",
		},
		{
			name: "with different types",
			entry: YamlLogEntry{
				Message: "different types",
				Extras: []any{
					"string_key", "string_value",
					"int_key", 42,
					"float_key", 3.14,
					"bool_key", true,
					"rune_key", 'X',
				},
			},
			expected: "msg: different types\nextras:\n  string_key: string_value\n  int_key: 42\n  float_key: 3.14\n  bool_key: true\n  rune_key: \"X\"\n",
		},
	}

	marshaler := NewYamlMarshaler()
	buf := bytes.NewBuffer(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshaler.MarshalInto(buf, tt.entry)
			result := buf.String()

			if result != tt.expected {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", tt.expected, result)
			}
			buf.Reset()
		})
	}
}

func TestYamlMarshaler_ContainsSpecialChars(t *testing.T) {
	yamlMarshaler := NewYamlMarshaler()
	tests := []struct {
		input    string
		expected bool
	}{
		{"simple", false},
		{"with:colon", true},
		{"with space", true},
		{"with-dash", true},
		{"with{brace}", true},
		{"with[bracket]", true},
		{"normal_underscore", false},
		{"normal.dot", false},
		{"with#hash", true},
		{"with@at", true},
		{"with!exclamation", true},
		{"with?question", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := yamlMarshaler.containsSpecialChars(tt.input)
			if result != tt.expected {
				t.Errorf("containsSpecialChars(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestYamlMarshaler_WriteValue(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		isKey    bool
		expected string
	}{
		{"string as key", "simple", true, "simple"},
		{"string as value", "simple", false, "simple"},
		{"string with special chars as key", "key:with:colon", true, "\"key:with:colon\""},
		{"string with special chars as value", "value:with:colon", false, "\"value:with:colon\""},
		{"integer as key", 42, true, "42"},
		{"integer as value", 42, false, "42"},
		{"float as key", 3.14, true, "3.14"},
		{"float as value", 3.14, false, "3.14"},
		{"boolean as key", true, true, "true"},
		{"boolean as value", false, false, "false"},
		{"rune as key", 'X', true, "X"},
		{"rune as value", 'Y', false, "\"Y\""},
	}

	marshaler := NewYamlMarshaler()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			marshaler.writeStr(&buf, tt.value, tt.isKey)
			if buf.String() != tt.expected {
				t.Errorf("writeStr(%v, %v) = %q, want %q",
					tt.value, tt.isKey, buf.String(), tt.expected)
			}
		})
	}
}
