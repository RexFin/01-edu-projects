package service

import (
	"ascii-art-web/pkg/hashs"
	"errors"
	"log"
	"os"
	"strings"
)

var expectedHashes = map[string]string{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52",
}

func GenerateAscii(text, banner string) (string, error) {
	content, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		log.Println(err)
		return "", err
	}

	isValid, err := hashs.CheckIntegrity(banner, expectedHashes)
	if err != nil || !isValid {
		log.Println("Error: banner file is corrupted")
		return "", errors.New("file corrupted")
	}

	rawLines := strings.Split(string(content), "\n")
	if banner == "thinkertoy" {
		rawLines = strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	}

	inputLines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	var output strings.Builder

	for _, word := range inputLines {
		if word == "" {
			output.WriteString("\n")
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, char := range word {
				start := int(char-32)*9 + i
				output.WriteString(rawLines[start])
			}
			output.WriteString("\n")
		}
	}
	return output.String(), nil
}
