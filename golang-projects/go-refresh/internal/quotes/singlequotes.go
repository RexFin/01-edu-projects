package quotes

import "strings"

func Singlequotes(slicedText []string) []string {
	isQuoteOpen := false
	openQuoteIdx := -1
	wordAfterOpenIdx := -1
	lastValidWordIdx := -1

	for i := 0; i < len(slicedText); i++ {
		word := slicedText[i]

		if word == "" {
			continue
		}

		if word == "'" {
			if isQuoteOpen {
				if openQuoteIdx != -1 && wordAfterOpenIdx != -1 {
					slicedText[wordAfterOpenIdx] = "'" + slicedText[wordAfterOpenIdx]
					slicedText[openQuoteIdx] = "" // удаляем одиночную
				}
				if lastValidWordIdx != -1 {
					slicedText[lastValidWordIdx] += "'"
				}
				slicedText[i] = ""

				// Сбрасываем состояние
				isQuoteOpen = false
				openQuoteIdx = -1
				wordAfterOpenIdx = -1
			} else {
				isQuoteOpen = true
				openQuoteIdx = i
				wordAfterOpenIdx = -1
			}
		} else if word != "\n" && word != "\r" {
			if isQuoteOpen && openQuoteIdx != -1 && wordAfterOpenIdx == -1 {
				wordAfterOpenIdx = i
			}
			if strings.HasPrefix(word, "'") && !strings.HasSuffix(word, "'") {
				isQuoteOpen = true
				openQuoteIdx = -1
			} else if strings.HasSuffix(word, "'") && !strings.HasPrefix(word, "'") && isQuoteOpen {
				if openQuoteIdx != -1 && wordAfterOpenIdx != -1 {
					slicedText[wordAfterOpenIdx] = "'" + slicedText[wordAfterOpenIdx]
					slicedText[openQuoteIdx] = ""
				}
				isQuoteOpen = false
				openQuoteIdx = -1
				wordAfterOpenIdx = -1
			}
			lastValidWordIdx = i
		}
	}

	return slicedText
}
