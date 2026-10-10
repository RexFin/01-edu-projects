package punct

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpacing(t *testing.T) {
	var tests = []struct {
		testName   string
		slicedText []string
		want       []string
	}{
		{
			testName:   "Just test 1",
			slicedText: []string{"hello", ","},
			want:       []string{"hello,", ""},
		},
		{
			testName:   "Just test 2",
			slicedText: []string{"wait", "...", "what", "?"},
			want:       []string{"wait...", "", "what?", ""},
		},
		{
			testName:   "Punctuation attached to word",
			slicedText: []string{"hello", ",world"},
			want:       []string{"hello,", "world"},
		},
		{
			testName:   "Multi punc is attached",
			slicedText: []string{"wait", "...yes!"},
			want:       []string{"wait...", "yes!"},
		},
		{
			testName:   "Ignore newline",
			slicedText: []string{"hello", "\n", ",world"},
			want:       []string{"hello", "\n", ",world"},
		},
		{
			testName:   "Double space between punct",
			slicedText: []string{"hello", "world", " ", " ", "!"},
			want:       []string{"hello", "world!", " ", " ", ""},
		},
		{
			testName:   "Nothing",
			slicedText: []string{"hello", "world,", "how", "are", "you?"},
			want:       []string{"hello", "world,", "how", "are", "you?"},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			ans := Spacing(tt.slicedText)
			if !reflect.DeepEqual(ans, tt.want) {
				fmt.Println("Test failed!")
				t.Errorf("[%s] got \n%s, want \n%s", tt.testName, ans, tt.want)
			}
		})
	}
}
