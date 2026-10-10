package pipeline

import (
	"go-refresh/internal/cases"
	"go-refresh/internal/grammar"
	"go-refresh/internal/numbers"
	"go-refresh/internal/punct"
	"go-refresh/internal/quotes"
)

func Pipeline(slicedText []string) []string {
	numbers.Convert(slicedText)
	cases.Transform(slicedText)

	slicedText = cleanEmptyElemFromSlice(slicedText)

	punct.Spacing(slicedText)
	quotes.Singlequotes(slicedText)
	grammar.Articles(slicedText)

	slicedText = cleanEmptyElemFromSlice(slicedText)

	return slicedText
}

func cleanEmptyElemFromSlice(slicedText []string) []string {
	j := 0
	for _, v := range slicedText {
		if v != "" {
			slicedText[j] = v
			j++
		}
	}

	for k := j; k < len(slicedText); k++ {
		slicedText[k] = ""
	}

	return slicedText[:j]
}
