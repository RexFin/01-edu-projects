package ai

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outChan := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outChan <- buf.String()
	}()

	f()
	w.Close()
	os.Stdout = oldStdout
	return <-outChan
}

func TestLangdetect(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "English text",
			input:    []string{"the", "cat", "is", "in", "the", "house"},
			expected: "Language: English (100%)\n",
		},
		{
			name:     "French text with dictionary words",
			input:    []string{"le", "chat", "est", "dans", "la", "maison"},
			expected: "Language: French (100%)\n",
		},
		{
			name:     "French text with special characters",
			input:    []string{"déjà", "vu", "café"},
			expected: "Language: French (100%)\n",
		},
		{
			name:     "Mixed text (50/50 - English priority)",
			input:    []string{"the", "le"},
			expected: "Language: English (50%)\n",
		},
		{
			name:     "Unknown language",
			input:    []string{"hello", "world", "privet"},
			expected: "Language: Unknown\n",
		},
		{
			name:     "Ignoring the cap cases (case-insensitive)",
			input:    []string{"The", "is.", "le,"},
			expected: "Language: Unknown\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				Langdetect(tt.input)
			})

			if output != tt.expected {
				t.Errorf("got %q, want %q", output, tt.expected)
			}
		})
	}
}
