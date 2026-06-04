package services

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteEscaped(t *testing.T) {
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
			WriteEscaped(buf, tt.input)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
