package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

var expectedHashes = map[string]string{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52",
}

func checkIntegrity(filename string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	actualHash := fmt.Sprintf("%x", h.Sum(nil))
	expectedHash, ok := expectedHashes[filename]
	if !ok {
		return false, fmt.Errorf("unknown banner")
	}

	return actualHash == expectedHash, nil
}

func draw(word string, alphabet []string) {
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			line := (int(char)-32)*9 + i
			if line < len(alphabet) {
				fmt.Print(alphabet[line])
			}
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 || args[0] == "" {
		return
	}

	input := args[0]
	if input == "\\n" {
		fmt.Println()
		return
	}

	banner := "standard.txt"
	if len(args) == 2 {
		banner = args[1] + ".txt"
	}

	isValid, err := checkIntegrity(banner)
	if err != nil || !isValid {
		fmt.Printf("Error: banner file %s is corrupted\n", banner)
		return
	}

	content, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error: could not find banner file", banner)
		return
	}

	alphabet := strings.Split(string(content), "\n")
	words := strings.Split(input, "\\n")

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		draw(word, alphabet)
	}
}
