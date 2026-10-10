package wordcount

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

func TestFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Basic test frequency",
			input:    []string{"apple", "banana", "apple", "orange", "banana", "apple", "orange", "kiwi", "grape", "melon"},
			expected: "Keywords: apple, banana, orange, grape, kiwi\n",
		},
		{
			name:     "Alphabetical sort (tie-breaker)",
			input:    []string{"zebra", "monkey", "apple", "dog", "cat"},
			expected: "Keywords: apple, cat, dog, monkey, zebra\n",
		},
		{
			name:     "Stopwords are ignored",
			input:    []string{"The", "cat", "is", "on", "the", "table", "it", "is", "a", "good", "cat"},
			expected: "Keywords: cat, good, table\n",
		},
		{
			name:     "Only stopwords",
			input:    []string{"the", "and", "of", "to", "is"},
			expected: "Keywords: (none)\n",
		},
		{
			name:     "Less than 5",
			input:    []string{"hello", "world", "hello"},
			expected: "Keywords: hello, world\n",
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: "Keywords: (none)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				Frequency(tt.input)
			})

			if output != tt.expected {
				t.Errorf("got %q, want %q", output, tt.expected)
			}
		})
	}
}
