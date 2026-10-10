package grammar

import (
	"strings"
	"unicode"
)

func Articles(slicedText []string) []string {
	for i := 0; i < len(slicedText)-1; i++ {
		current := slicedText[i]

		if current != "a" && current != "A" && current != "an" && current != "An" {
			continue
		}

		nextWord := slicedText[i+1]

		if len(nextWord) == 0 {
			continue
		}

		firstChar := unicode.ToLower(rune(nextWord[0]))
		isVowelOrH := strings.ContainsRune("aeiouh", firstChar)
		isUpper := current == "A" || current == "An"

		if isVowelOrH {
			if isUpper {
				slicedText[i] = "An"
			} else {
				slicedText[i] = "an"
			}
		} else {
			if isUpper {
				slicedText[i] = "A"
			} else {
				slicedText[i] = "a"
			}
		}
	}

	return slicedText
}
