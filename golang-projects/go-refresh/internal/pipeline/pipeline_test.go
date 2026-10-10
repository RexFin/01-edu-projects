package pipeline

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCleanEmptyElemFromSlice(t *testing.T) {
	var tests = []struct {
		testName   string
		slicedText []string
		want       []string
	}{
		{
			testName:   "Mixed test",
			slicedText: []string{"hello", "", "world", "", "", "!"},
			want:       []string{"hello", "world", "!"},
		},
		{
			testName:   "No empty strings",
			slicedText: []string{"just", "normal", "words"},
			want:       []string{"just", "normal", "words"},
		},
		{
			testName:   "Only empty",
			slicedText: []string{"", "", ""},
			want:       []string{},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			inputCopy := make([]string, len(tt.slicedText))
			copy(inputCopy, tt.slicedText)

			ans := cleanEmptyElemFromSlice(inputCopy)

			if !reflect.DeepEqual(ans, tt.want) {
				fmt.Println("Test failed!")
				t.Errorf("[%s] got \n%q, want \n%q", tt.testName, ans, tt.want)
			}
		})
	}
}
