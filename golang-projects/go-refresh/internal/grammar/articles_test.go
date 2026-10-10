package grammar

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArticles(t *testing.T) {
	var tests = []struct {
		testName   string
		slicedText []string
		want       []string
	}{
		{
			testName:   "First easy test",
			slicedText: []string{"there", "is", "a", "apple"},
			want:       []string{"there", "is", "an", "apple"},
		},
		{
			testName:   "Second easy test",
			slicedText: []string{"it", "takes", "a", "hour"},
			want:       []string{"it", "takes", "an", "hour"},
		},
		{
			testName:   "Third test",
			slicedText: []string{"what", "an", "beautiful", "day"},
			want:       []string{"what", "a", "beautiful", "day"},
		},
		{
			testName:   "Cap test",
			slicedText: []string{"A", "elephant"},
			want:       []string{"An", "elephant"},
		},
		{
			testName:   "Cap test 2",
			slicedText: []string{"An", "Cat"},
			want:       []string{"A", "Cat"},
		},
		{
			testName:   "Next word is cap",
			slicedText: []string{"eat", "a", "Orange"},
			want:       []string{"eat", "an", "Orange"},
		},
		{
			testName:   "Just a prank",
			slicedText: []string{"a", "dog", "and", "an", "owl"},
			want:       []string{"a", "dog", "and", "an", "owl"},
		},
		{
			testName:   "Out of bound test",
			slicedText: []string{"give", "me", "a"},
			want:       []string{"give", "me", "a"},
		},
		{
			testName:   "Empty test",
			slicedText: []string{"a", "", "apple"},
			want:       []string{"a", "", "apple"},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			inputCopy := make([]string, len(tt.slicedText))
			copy(inputCopy, tt.slicedText)

			ans := Articles(inputCopy)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("\n[%s] failed!\ngot:  %q\nwant: %q", tt.testName, ans, tt.want)
			}
		})
	}
}
