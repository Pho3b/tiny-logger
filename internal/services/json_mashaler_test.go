package services

import (
	"encoding/json"
	"github.com/pho3b/tiny-logger/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMarshaler_Marshal_MessageOnly(t *testing.T) {
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Message: "basic message",
	}

	got := m.Marshal(entry)
	want := `{"msg":"basic message"}`
	if string(got) != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_SomeFields(t *testing.T) {
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:   "warn",
		Message: "something odd happened",
		Time:    "12:34:56",
	}

	got := m.Marshal(entry)
	want := `{"level":"warn","time":"12:34:56","msg":"something odd happened"}`
	if string(got) != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_AllFields(t *testing.T) {
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:    "info",
		Date:     "2025-06-14",
		Time:     "20:15:30",
		DateTime: "2025-06-14T20:15:30",
		Message:  "all systems go",
	}

	got := m.Marshal(entry)
	want := `{"level":"info","date":"2025-06-14","time":"20:15:30","datetime":"2025-06-14T20:15:30","msg":"all systems go"}`
	if string(got) != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_WithExtras(t *testing.T) {
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:    "INFO",
		Date:     "",
		Time:     "",
		DateTime: "20/06/2025 08:11:06",
		Message:  "all systems go",
		Extras:   []any{"bool", true, "int", 3, "float", 4.3, "arr", []int{1, 2, 3}, "rune", 'A', "string", "ciaooo", "null"},
	}

	got := m.Marshal(entry)
	want := "{\"level\":\"INFO\",\"datetime\":\"20/06/2025 08:11:06\",\"msg\":\"all systems go\"," +
		"\"extras\":{\"bool\":true,\"int\":3,\"float\":4.3,\"arr\":\"[1 2 3]\",\"rune\":\"A\",\"string\":\"ciaooo\",\"null\":null}}"
	if string(got) != want {
		t.Errorf("got = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Unmarshal_Std_Marshal_Result(t *testing.T) {
	m := JsonMarshaler{}
	entry := JsonLogEntry{
		Level:    "INFO",
		Date:     "",
		Time:     "",
		DateTime: "20/06/2025 08:11:06",
		Message:  "all systems go",
		Extras:   []any{"bool", true, "int", 3, "float", 4.3, "arr", []int{1, 2, 3}, "rune", 'A', "string", "ciaooo", "null"},
	}

	jsonMsg := m.Marshal(entry)
	assert.NoError(t, json.Unmarshal(jsonMsg, &shared.JsonLog{}))
}
