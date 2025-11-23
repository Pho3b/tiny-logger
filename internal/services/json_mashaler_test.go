package services

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/pho3b/tiny-logger/shared"
	"github.com/stretchr/testify/assert"
)

func TestJsonMarshaler_Marshal_MessageOnly(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Message: "basic message",
	}

	m.MarshalInto(buf, entry)
	want := `{"msg":"basic message"}`
	got := buf.String()
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_SomeFields(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:   "warn",
		Message: "something odd happened",
		Time:    "12:34:56",
		Extras:  []any{"rune", ':', "int", 23},
	}

	m.MarshalInto(buf, entry)
	got := buf.String()
	want := `{"level":"warn","time":"12:34:56","msg":"something odd happened","extras":{"rune":":","int":23}}`
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_AllFields(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:          "info",
		Date:           "2025-06-14",
		Time:           "20:15:30",
		DateTime:       "2025-06-14T20:15:30",
		Message:        "all systems go",
		DateTimeFormat: shared.IT,
	}

	m.MarshalInto(buf, entry)
	got := buf.String()
	want := `{"level":"info","datetime":"2025-06-14T20:15:30","msg":"all systems go"}`
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_WithExtras(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:    "INFO",
		Date:     "",
		Time:     "",
		DateTime: "20/06/2025 08:11:06",
		Message:  "all systems go",
		Extras:   []any{"bool", true, "int", 3, "float", 4.3, "arr", []int{1, 2, 3}, "rune", 'A', "string", "ciaooo", "null"},
	}

	m.MarshalInto(buf, entry)
	got := buf.String()
	want := "{\"level\":\"INFO\",\"datetime\":\"20/06/2025 08:11:06\",\"msg\":\"all systems go\"," +
		"\"extras\":{\"bool\":true,\"int\":3,\"float\":4.3,\"arr\":\"[1 2 3]\",\"rune\":\"A\",\"string\":\"ciaooo\",\"null\":null}}"
	if got != want {
		t.Errorf("got = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Unmarshal_Std_Marshal_Result(t *testing.T) {
	buf := &bytes.Buffer{}
	m := JsonMarshaler{}
	entry := JsonLogEntry{
		Level:    "INFO",
		Date:     "",
		Time:     "",
		DateTime: "20/06/2025 08:11:06",
		Message:  "all systems go",
		Extras:   []any{"bool", true, "int", 3, "float", 4.3, "arr", []int{1, 2, 3}, "rune", 'A', "string", "ciaooo", "null"},
	}

	m.MarshalInto(buf, entry)
	assert.NoError(t, json.Unmarshal(buf.Bytes(), &shared.JsonLog{}))
}

func TestJsonMarshaler_Marshal_UnixTimestamp(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:          "info",
		DateTime:       "1700000000",
		Message:        "unix time",
		DateTimeFormat: shared.UnixTimestamp,
	}

	m.MarshalInto(buf, entry)
	got := buf.String()
	// Expecting "ts" key instead of "datetime"
	want := `{"level":"info","ts":"1700000000","msg":"unix time"}`
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_OnlyDate(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:   "info",
		Date:    "23/10/2022",
		Message: "only date",
	}

	m.MarshalInto(buf, entry)
	got := buf.String()
	// Expecting "ts" key instead of "datetime"
	want := `{"level":"info","date":"23/10/2022","msg":"only date"}`
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_OnlyTime(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:   "info",
		Time:    "16:00",
		Message: "only time",
	}

	m.MarshalInto(buf, entry)
	got := buf.String()
	// Expecting "ts" key instead of "datetime"
	want := `{"level":"info","time":"16:00","msg":"only time"}`
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}
