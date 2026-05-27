package pkg

import (
	"os"
	"strings"
)

func GenerateAscii(text, banner string) (string, error) {
	content, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return "", err
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
