package services

import (
	"encoding/json"
	"testing"
)

func TestJsonMarshaler_Marshal_MessageOnly(t *testing.T) {
	m := &JsonMarshaler{}
	entry := &jsonLogEntry{
		Message: "basic message",
	}

	got, err := m.Marshal(entry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := `{"msg":"basic message"}`
	if string(got) != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_SomeFields(t *testing.T) {
	m := &JsonMarshaler{}
	entry := &jsonLogEntry{
		Level:   "warn",
		Message: "something odd happened",
		Time:    "12:34:56",
	}

	got, err := m.Marshal(entry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := `{"level":"warn","time":"12:34:56","msg":"something odd happened"}`
	if string(got) != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_AllFields(t *testing.T) {
	m := &JsonMarshaler{}
	entry := &jsonLogEntry{
		Level:    "info",
		Date:     "2025-06-14",
		Time:     "20:15:30",
		DateTime: "2025-06-14T20:15:30",
		Message:  "all systems go",
	}

	got, err := m.Marshal(entry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := `{"level":"info","date":"2025-06-14","time":"20:15:30","datetime":"2025-06-14T20:15:30","msg":"all systems go"}`
	if string(got) != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_WithExtras(t *testing.T) {
	m := &JsonMarshaler{}
	entry := &jsonLogEntry{
		Level:    "INFO",
		Date:     "",
		Time:     "",
		DateTime: "2025-06-14T20:15:30",
		Message:  "all systems go",
		Extras:   []any{"bool", true, "int", 3, "float", 4.3, "arr", []int{1, 2, 3}, "rune", 'A', "string", "ciaooo", "null"},
	}

	got, err := m.Marshal(entry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want, _ := json.Marshal(entry)
	if string(got) != string(want) {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}
