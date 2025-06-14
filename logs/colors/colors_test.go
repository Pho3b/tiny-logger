package colors

import (
	"runtime"
	"strings"
	"testing"
)

func TestColorString(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{"Reset", Reset, "\033[0m"},
		{"Red", Red, "\033[31m"},
		{"Magenta", Magenta, "\033[91m"},
		{"Yellow", Yellow, "\033[33m"},
		{"Cyan", Cyan, "\033[36m"},
		{"Gray", Gray, "\033[37m"},
		{"Black", Black, "\033[30m"},
		{"Green", Green, "\033[32m"},
		{"Blue", Blue, "\033[34m"},
		{"White", White, "\033[97m"},
		{"DarkGray", DarkGray, "\033[90m"},
		{"BrightGreen", BrightGreen, "\033[92m"},
		{"BrightYellow", BrightYellow, "\033[93m"},
		{"BrightBlue", BrightBlue, "\033[94m"},
		{"Pink", Pink, "\033[95m"},
		{"BrightCyan", BrightCyan, "\033[96m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colorCode := tt.color.String()

			// On Windows, expect empty string
			if isWindows() {
				if colorCode != "" {
					t.Errorf("Color %s: expected empty string on Windows, got %q", tt.name, colorCode)
				}
			} else {
				if colorCode != tt.expected {
					t.Errorf("Color %s: expected %q, got %q", tt.name, tt.expected, colorCode)
				}
			}
		})
	}
}

// isWindows checks if the runtime OS is Windows
func isWindows() bool {
	return strings.Contains(runtime.GOOS, "windows")
}
