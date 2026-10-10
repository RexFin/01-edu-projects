package modfields

func ModificationFix(rawText []rune) []string {
	slicedText := make([]string, 0, len(rawText)/5)
	var currentWord []rune

	for _, v := range rawText {
		switch v {
		case ' ', '\t':
			if len(currentWord) > 0 {
				slicedText = append(slicedText, string(currentWord))
				currentWord = nil
			}
		case '\n', '\r':
			if len(currentWord) > 0 {
				slicedText = append(slicedText, string(currentWord))
				currentWord = nil
			}
			slicedText = append(slicedText, string(v))
		case '(':
			if len(currentWord) > 0 {
				slicedText = append(slicedText, string(currentWord))
				currentWord = nil
			}
			currentWord = append(currentWord, v)
		case ')':
			currentWord = append(currentWord, v)
			slicedText = append(slicedText, string(currentWord))
			currentWord = nil
		default:
			currentWord = append(currentWord, v)
		}
	}

	if len(currentWord) > 0 {
		slicedText = append(slicedText, string(currentWord))
	}

	return slicedText
}
