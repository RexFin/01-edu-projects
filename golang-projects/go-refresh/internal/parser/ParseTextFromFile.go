package parser

import (
	"fmt"
	"os"
)

func ParseTextFromFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	result := make([]byte, 0, len(data))

	lastWasSpace := false
	isStartOfLine := true

	for _, b := range data {
		switch b {
		case ' ', '\t':
			if !isStartOfLine && !lastWasSpace {
				result = append(result, ' ')
				lastWasSpace = true
			}
		case '\n', '\r':
			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}
			result = append(result, b)
			isStartOfLine = true
			lastWasSpace = false
		default:
			result = append(result, b)
			isStartOfLine = false
			lastWasSpace = false
		}
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result, nil
}
