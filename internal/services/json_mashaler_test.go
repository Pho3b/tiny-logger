package services

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Pho3b/tiny-logger/shared"
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
		Level:   "info",
		Date:    "2025-06-14",
		Time:    "20:15:30",
		Message: "all systems go",
		UnixTS:  "",
	}

	m.MarshalInto(buf, entry)
	got := buf.String()
	want := `{"level":"info","datetime":"2025-06-14 20:15:30","msg":"all systems go"}`
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestJsonMarshaler_Marshal_WithExtras(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:   "INFO",
		Date:    "20/06/2025",
		Time:    "08:11:06",
		Message: "all systems go",
		Extras:  []any{"bool", true, "int", 3, "float", 4.3, "arr", []int{1, 2, 3}, "rune", 'A', "string", "ciaooo", "null"},
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
		Level:   "INFO",
		Date:    "20/06/2025",
		Time:    "08:11:06",
		Message: "all systems go",
		Extras:  []any{"bool", true, "int", 3, "float", 4.3, "arr", []int{1, 2, 3}, "rune", 'A', "string", "ciaooo", "null"},
	}

	m.MarshalInto(buf, entry)
	assert.NoError(t, json.Unmarshal(buf.Bytes(), &shared.JsonLog{}))
}

func TestJsonMarshaler_Marshal_UnixTimestamp(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:   "info",
		Message: "unix time",
		UnixTS:  "1700000000",
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

func TestJsonMarshaler_Marshal_EscapesSpecialChars(t *testing.T) {
	buf := &bytes.Buffer{}
	m := &JsonMarshaler{}
	entry := JsonLogEntry{
		Level:   "INFO",
		Message: "he said \"hi\"\n\tand left\\done",
		Extras:  []any{"quote\"key", "a\"b\\c\nd"},
	}

	m.MarshalInto(buf, entry)
	got := buf.String()

	// The produced output must be valid JSON that round-trips the original values.
	var decoded shared.JsonLog
	assert.NoError(t, json.Unmarshal(buf.Bytes(), &decoded), "output is not valid JSON: %s", got)
	assert.Equal(t, "he said \"hi\"\n\tand left\\done", decoded.Message)
	assert.Equal(t, "a\"b\\c\nd", decoded.Extras["quote\"key"])
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
	want := `{"level":"info","time":"16:00","msg":"only time"}`
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}

func TestWriteJSONEscaped(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "clean string takes fast path",
			input: "hello world",
			want:  "hello world",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "double quote escaped",
			input: `say "hi"`,
			want:  `say \"hi\"`,
		},
		{
			name:  "backslash escaped",
			input: `a\b`,
			want:  `a\\b`,
		},
		{
			name:  "newline escaped",
			input: "line1\nline2",
			want:  `line1\nline2`,
		},
		{
			name:  "carriage return escaped",
			input: "a\rb",
			want:  `a\rb`,
		},
		{
			name:  "tab escaped",
			input: "a\tb",
			want:  `a\tb`,
		},
		{
			name:  "control char 0x01 rendered as unicode escape",
			input: "\x01",
			want:  "\\u0001",
		},
		{
			name:  "control char 0x1b (ESC) rendered as unicode escape",
			input: "\x1b",
			want:  "\\u001b",
		},
		{
			name:  "consecutive escapable chars skip empty flush",
			input: "\n\n",
			want:  `\n\n`,
		},
		{
			name:  "escape in the middle preserves clean prefix and suffix",
			input: "prefix\nsuffix",
			want:  `prefix\nsuffix`,
		},
		{
			name:  "multiple different escapes in one string",
			input: "a\"b\\c\nd\te",
			want:  `a\"b\\c\nd\te`,
		},
		{
			name:  "string ending with escapable char flushes tail correctly",
			input: "abc\"",
			want:  `abc\"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			writeJSONEscaped(buf, tt.input)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
