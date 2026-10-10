package modfields

import (
	"fmt"
	"reflect"
	"testing"
)

func TestModificationFix(t *testing.T) {
	var tests = []struct {
		testName string
		rawText  []rune
		want     []string
	}{
		{
			testName: "Normal words separation by space and tabs",
			rawText:  []rune("Hello world \t test"),
			want:     []string{"Hello", "world", "test"},
		},
		{
			testName: "Newlines and carriage returns preservation",
			rawText:  []rune("Hello\nworld\r\nGo"),
			want:     []string{"Hello", "\n", "world", "\r", "\n", "Go"},
		},
		{
			testName: "Single modifier in brackets",
			rawText:  []rune("1E (hex) is 30"),
			want:     []string{"1E", "(hex)", "is", "30"},
		},
		{
			testName: "Parentheses attached to words without spaces",
			rawText:  []rune("word(up)another"),
			want:     []string{"word", "(up)", "another"},
		},
		{
			testName: "Empty input",
			rawText:  []rune(""),
			want:     []string{},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			ans := ModificationFix(tt.rawText)
			if len(ans) == 0 && len(tt.want) == 0 {
				ans = []string{}
			}
			if !reflect.DeepEqual(ans, tt.want) {
				fmt.Println("Test failed!")
				t.Errorf("[%s] got \n%q, want \n%q", tt.testName, ans, tt.want)
			}
		})
	}
}
