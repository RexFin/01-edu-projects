package wordcount

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func Frequency(slicedText []string) {
	stopwords := map[string]struct{}{
		"the": {}, "and": {}, "of": {}, "to": {}, "is": {}, "in": {}, "it": {}, "that": {},
		"a": {}, "an": {}, "for": {}, "on": {}, "with": {}, "as": {}, "by": {}, "at": {},
		"from": {}, "this": {}, "these": {}, "those": {}, "de": {}, "la": {}, "le": {},
		"les": {}, "et": {}, "un": {}, "une": {}, "des": {},
	}

	freqMap := make(map[string]int)

	for _, text := range slicedText {
		text = strings.ToLower(text)

		words := strings.FieldsFunc(text, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
		})

		for _, w := range words {
			w = strings.Trim(w, "'")

			if w == "" {
				continue
			}
			if _, isStopword := stopwords[w]; !isStopword {
				freqMap[w]++
			}
		}
	}

	if len(freqMap) == 0 {
		fmt.Println("Keywords: (none)")
		return
	}

	type wordFreq struct {
		word  string
		count int
	}

	wordsList := make([]wordFreq, 0, len(freqMap))
	for w, c := range freqMap {
		wordsList = append(wordsList, wordFreq{word: w, count: c})
	}

	sort.Slice(wordsList, func(i, j int) bool {
		if wordsList[i].count == wordsList[j].count {
			return wordsList[i].word < wordsList[j].word
		}
		return wordsList[i].count > wordsList[j].count
	})

	limit := 5
	if len(wordsList) < limit {
		limit = len(wordsList)
	}

	topWords := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		topWords = append(topWords, wordsList[i].word)
	}

	fmt.Printf("Keywords: %s\n", strings.Join(topWords, ", "))
}
