package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseTextFromFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "parser_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	var tests = []struct {
		testName    string
		fileContent string
		want        []byte
		wantError   bool
		nonExistent bool
	}{
		{
			testName:    "Normal text",
			fileContent: "Hello world\nThis is a test.",
			want:        []byte("Hello world\nThis is a test."),
			wantError:   false,
		},
		{
			testName:    "Multiple spaces and tabs",
			fileContent: "Hello    world\t\tThis is\ta test.",
			want:        []byte("Hello world This is a test."),
			wantError:   false,
		},
		{
			testName:    "Leading spaces and tabs on lines",
			fileContent: "   Hello world\n\tThis is a test.",
			want:        []byte("Hello world\nThis is a test."),
			wantError:   false,
		},
		{
			testName:    "Trailing spaces on lines",
			fileContent: "Hello world   \nThis is a test. \t",
			want:        []byte("Hello world\nThis is a test."),
			wantError:   false,
		},
		{
			testName:    "Empty file",
			fileContent: "",
			want:        []byte(""),
			wantError:   false,
		},
		{
			testName:    "File does not exist",
			fileContent: "",
			want:        nil,
			wantError:   true,
			nonExistent: true,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			var filePath string

			if tt.nonExistent {
				filePath = filepath.Join(tempDir, "does_not_exist.txt")
			} else {
				file, _ := os.CreateTemp(tempDir, "test_*.txt")
				file.WriteString(tt.fileContent)
				file.Close()
				filePath = file.Name()
			}

			ans, err := ParseTextFromFile(filePath)

			if (err != nil) != tt.wantError {
				fmt.Println("Test failed!")
				t.Errorf("[%s] got error \n%v, want error \n%v", tt.testName, err != nil, tt.wantError)
				return
			}

			if !tt.wantError && !reflect.DeepEqual(ans, tt.want) {
				fmt.Println("Test failed!")
				t.Errorf("[%s] got \n%s, want \n%s", tt.testName, string(ans), string(tt.want))
			}
		})
	}
}
