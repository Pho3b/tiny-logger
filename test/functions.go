package test

import (
	"bytes"
	"os"
)

// CaptureOutput redirects os.Stdout to capture the output of the function f
func CaptureOutput(f func()) string {
	r, w, _ := os.Pipe()
	defer r.Close()

	origStdout := os.Stdout
	os.Stdout = w

	f()
	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

// CaptureErrorOutput redirects os.Stderr to capture the output of the function f
func CaptureErrorOutput(f func()) string {
	r, w, _ := os.Pipe()
	defer r.Close()

	origStderr := os.Stderr
	os.Stderr = w

	f()
	w.Close()
	os.Stderr = origStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
