package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		expectErr bool
	}{
		{
			name: "Vaild test 1",
			content: "....\n" +
				".##.\n" +
				".##.\n" +
				"....",
			expectErr: false,
		},
		{
			name: "Valid test 2",
			content: "....\n" +
				"####\n" +
				"....\n" +
				"....\n" +
				"\n" +
				"...#\n" +
				"...#\n" +
				"...#\n" +
				"...#",
			expectErr: false,
		},
		{
			name: "Windows is trash",
			content: "....\r\n" +
				"####\r\n" +
				"....\r\n" +
				"....\r\n" +
				"\r\n" +
				"##..\r\n" +
				"##..\r\n" +
				"....\r\n" +
				"....",
			expectErr: false,
		},
		{
			name: "Not a block character",
			content: "....\n" +
				".##.\n" +
				".#A.\n" +
				"....",
			expectErr: true,
		},
		{
			name: "Less than 4 blocks",
			content: "....\n" +
				".##.\n" +
				".#..\n" +
				"....",
			expectErr: true,
		},
		{
			name: "More than 4 blocks",
			content: "....\n" +
				".##.\n" +
				"###.\n" +
				"....",
			expectErr: true,
		},
		{
			name: "Disconnected shape",
			content: "#..#\n" +
				"....\n" +
				"....\n" +
				"#..#",
			expectErr: true,
		},
		{
			name: "Invalid width",
			content: ".....\n" +
				"####.\n" +
				".....\n" +
				".....",
			expectErr: true,
		},
		{
			name: "Invalid height",
			content: "....\n" +
				"####\n" +
				"....",
			expectErr: true,
		},
		{
			name: "Invalid separator",
			content: "....\n" +
				"####\n" +
				"....\n" +
				"....\n" +
				" \n" +
				"...#\n" +
				"...#\n" +
				"...#\n" +
				"...#",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			path := filepath.Join(tempDir, "test.txt")

			err := os.WriteFile(path, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}

			_, err = Parse(path)
			hasErr := err != nil

			if hasErr != tt.expectErr {
				t.Errorf("\nwanted err: %v\nget err:  %v", tt.expectErr, err)
			}
		})
	}
}
