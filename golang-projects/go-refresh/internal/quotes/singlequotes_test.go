package quotes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSinglequotes(t *testing.T) {
	var tests = []struct {
		testName   string
		slicedText []string
		want       []string
	}{
		{
			testName:   "First test",
			slicedText: []string{"'", "hello", "'"},
			want:       []string{"", "'hello'", ""},
		},
		{
			testName:   "Second test",
			slicedText: []string{"'", "hello", "world", "'"},
			want:       []string{"", "'hello", "world'", ""},
		},
		{
			testName:   "No second quote",
			slicedText: []string{"'", "hello", "world"},
			want:       []string{"'", "hello", "world"},
		},
		{
			testName:   "Quete is attached text",
			slicedText: []string{"'hello", "'"},
			want:       []string{"'hello'", ""},
		},
		{
			testName:   "Quete is attached text v2",
			slicedText: []string{"'", "world'"},
			want:       []string{"", "'world'"},
		},
		{
			testName:   "Ignore empty",
			slicedText: []string{"'", "", "hello", "", "'"},
			want:       []string{"", "", "'hello'", "", ""},
		},
		{
			testName:   "Skip newline",
			slicedText: []string{"'", "\n", "hello", "world", "\n", "'"},
			want:       []string{"", "\n", "'hello", "world'", "\n", ""},
		},
		{
			testName:   "Inner quote is ignored",
			slicedText: []string{"please", "don't", "touch", "it's", "mine"},
			want:       []string{"please", "don't", "touch", "it's", "mine"},
		},
		{
			testName:   "Quote in quote in quote etc...",
			slicedText: []string{"'", "one", "'", "and", "'", "two", "'"},
			want:       []string{"", "'one'", "", "and", "", "'two'", ""},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			ans := Singlequotes(tt.slicedText)
			if !reflect.DeepEqual(ans, tt.want) {
				fmt.Println("Test failed!")
				t.Errorf("[%s] got \n%s, want \n%s", tt.testName, ans, tt.want)
			}
		})
	}
}
