package cases

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTransform(t *testing.T) {
	var tests = []struct {
		testName   string
		slicedText []string
		want       []string
	}{
		{
			testName:   "First easy test",
			slicedText: []string{"hello", "world", "(up)"},
			want:       []string{"hello", "WORLD", ""},
		},
		{
			testName:   "Second easy test",
			slicedText: []string{"TWO", "NUMBER", "NINE", "(low)"},
			want:       []string{"TWO", "NUMBER", "nine", ""},
		},
		{
			testName:   "Third easy test",
			slicedText: []string{"gotta", "go", "fAst", "(cap)"},
			want:       []string{"gotta", "go", "Fast", ""},
		},

		{
			testName:   "Fourth easy test",
			slicedText: []string{"make", "these", "words", "big", "(up,", "3)"},
			want:       []string{"make", "THESE", "WORDS", "BIG", "", ""},
		},
		{
			testName:   "Fifth easy test",
			slicedText: []string{"hello", "beautiful", "world", "(cap,", "2)"},
			want:       []string{"hello", "Beautiful", "World", "", ""},
		},

		{
			testName:   "Number bigger than words",
			slicedText: []string{"one", "two", "(up,", "5)"},
			want:       []string{"ONE", "TWO", "", ""},
		},
		{
			testName:   "Skip newlines",
			slicedText: []string{"hello", "\n", "world", "(cap,", "2)"},
			want:       []string{"Hello", "\n", "World", "", ""},
		},
		{
			testName:   "Tag spam",
			slicedText: []string{"hello", "world", "(up)", "(low)"},
			want:       []string{"hello", "world", "", ""},
		},
		{
			testName:   "No number in multi up",
			slicedText: []string{"hello", "world", "(up,"},
			want:       []string{"hello", "world", ""},
		},
		{
			testName:   "Try to apply the correct word",
			slicedText: []string{"hello", "world", "!", "(up,", "2)"},
			want:       []string{"HELLO", "WORLD", "!", "", ""},
		},
		{
			testName:   "Ignore negative numbers",
			slicedText: []string{"hello", "world", "(up,", "-2)"},
			want:       []string{"hello", "world", "", ""},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			inputCopy := make([]string, len(tt.slicedText))
			copy(inputCopy, tt.slicedText)

			ans := Transform(inputCopy)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("\n[%s] failed!\ngot:  %q\nwant: %q", tt.testName, ans, tt.want)
			}
		})
	}
}
