package cases

import (
	"strconv"
	"strings"
	"unicode"
)

func Transform(slicedText []string) []string {
	for i := 0; i < len(slicedText); i++ {
		var action func(string) string

		count := 1
		isMulti := false

		switch strings.ToLower(slicedText[i]) {
		case "(up)":
			action = strings.ToUpper
		case "(low)":
			action = strings.ToLower
		case "(cap)":
			action = capitalize
		case "(up,":
			action = strings.ToUpper
			isMulti = true
		case "(low,":
			action = strings.ToLower
			isMulti = true
		case "(cap,":
			action = capitalize
			isMulti = true
		default:
			continue
		}

		if isMulti {
			if i+1 < len(slicedText) {
				count = extractNumber(slicedText[i+1])
				slicedText[i+1] = ""
			} else {
				count = 0
			}
		}

		applied := 0

		for j := i - 1; j >= 0 && applied < count; j-- {
			if isValidWord(slicedText[j]) {
				slicedText[j] = action(slicedText[j])
				applied++
			}
		}

		slicedText[i] = ""
	}

	return slicedText
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	foundFirst := false

	for i, r := range runes {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if !foundFirst {
				runes[i] = unicode.ToUpper(r)
				foundFirst = true
			} else {
				runes[i] = unicode.ToLower(r)
			}
		}
	}

	return string(runes)
}

func extractNumber(s string) int {
	var numStr strings.Builder

	for _, r := range s {
		if unicode.IsDigit(r) {
			numStr.WriteRune(r)
		} else if r == '-' && numStr.Len() == 0 {
			numStr.WriteRune(r)
		}
	}

	num, err := strconv.Atoi(numStr.String())
	if err != nil {
		return 0
	}

	if num < 0 {
		return 0
	}

	return num
}

func isValidWord(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
