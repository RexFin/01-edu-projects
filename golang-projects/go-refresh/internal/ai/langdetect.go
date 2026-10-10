package ai

import (
	"fmt"
	"strings"
)

func Langdetect(slicedText []string) {
	engCount := 0
	frCount := 0

	engWordSet := map[string]struct{}{
		"the": {}, "and": {}, "of": {}, "to": {}, "is": {},
		"in": {}, "you": {}, "it": {}, "that": {},
	}

	frnWordSet := map[string]struct{}{
		"le": {}, "la": {}, "les": {}, "de": {}, "et": {},
		"un": {}, "une": {}, "est": {}, "dans": {}, "vous": {},
	}

	for _, word := range slicedText {
		if _, ok := engWordSet[word]; ok {
			engCount++
		}

		if _, ok := frnWordSet[word]; ok {
			frCount++
		}

		if strings.ContainsAny(word, "éèàçôîû") {
			frCount++
		}
	}

	total := engCount + frCount

	if total == 0 {
		fmt.Println("Language: Unknown")
		return
	}

	if engCount >= frCount {
		percent := (engCount*100 + total/2) / total
		fmt.Printf("Language: English (%d%%)\n", percent)
	} else {
		percent := (frCount*100 + total/2) / total
		fmt.Printf("Language: French (%d%%)\n", percent)
	}
}
