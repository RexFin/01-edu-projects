package numbers

import (
	"strconv"
	"strings"
)

func Convert(slicedText []string) []string {
	for i := 0; i < len(slicedText); i++ {
		var convertMethod func(string) string

		switch strings.ToLower(slicedText[i]) {
		case "(hex)":
			convertMethod = hexToDecimal
		case "(bin)":
			convertMethod = binToDecimal
		default:
			continue
		}

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
			slicedText[targetIdx] = convertMethod(slicedText[targetIdx])
		}

		slicedText[i] = ""
	}

	return slicedText
}

func hexToDecimal(hexStr string) string {
	cleaned := strings.TrimPrefix(hexStr, "0x")
	cleaned = strings.TrimPrefix(cleaned, "0X")

	result, err := strconv.ParseUint(cleaned, 16, 64)
	if err != nil {
		return hexStr
	}
	return strconv.FormatUint(result, 10)
}

func binToDecimal(bin string) string {
	var result int64 = 0

	for _, c := range bin {
		if c != '0' && c != '1' {
			return bin
		}

		result = result*2 + int64(c-'0')
	}

	return strconv.FormatInt(result, 10)
}
