package punct

import (
	"strings"
)

func Spacing(slicedText []string) []string {
	for i := 0; i < len(slicedText); i++ {
		word := slicedText[i]

		if word == "" {
			continue
		}

		runes := []rune(word)
		puncCount := 0

		for _, r := range runes {
			if strings.ContainsRune(".,!?:;", r) {
				puncCount++
			} else {
				break
			}
		}

		if puncCount > 0 {
			targetIdx := -1
			for j := i - 1; j >= 0; j-- {
				if slicedText[j] == "" || slicedText[j] == " " || slicedText[j] == "\t" {
					continue
				}
				if slicedText[j] == "\n" || slicedText[j] == "\r" {
					break
				}
				targetIdx = j
				break
			}

			if targetIdx != -1 {
				puncs := string(runes[:puncCount])
				restOfWord := string(runes[puncCount:])

				slicedText[targetIdx] += puncs
				slicedText[i] = restOfWord
			}
		}
	}

	return slicedText
}
