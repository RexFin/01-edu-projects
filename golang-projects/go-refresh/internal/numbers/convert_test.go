package numbers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	var tests = []struct {
		testName   string
		slicedText []string
		want       []string
	}{
		{
			testName:   "First normal test",
			slicedText: []string{"5", "+", "5", "is", "A", "(hex)"},
			want:       []string{"5", "+", "5", "is", "10", ""},
		},
		{
			testName:   "Second normal test",
			slicedText: []string{"101010", "(bin)", "BRATUHA", "!"},
			want:       []string{"42", "", "BRATUHA", "!"},
		},
		{
			testName:   "Hex spam test",
			slicedText: []string{"(hex)", "(hex)", "(hex)", "(hex)", "16", "(hex)"},
			want:       []string{"", "", "", "", "22", ""},
		},
		{
			testName:   "Bin error test",
			slicedText: []string{"(hex)", "(hex)", "(hex)", "(hex)", "10IAMNOTABIN10", "(bin)"},
			want:       []string{"", "", "", "", "10IAMNOTABIN10", ""},
		},
		{
			testName:   "Double modificator test",
			slicedText: []string{"a", "(hex)", "(bin)"},
			want:       []string{"2", "", ""},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			ans := Convert(tt.slicedText)
			if !reflect.DeepEqual(ans, tt.want) {
				fmt.Println("Test failed!")
				t.Errorf("[%s] got \n%s, want \n%s", tt.testName, ans, tt.want)
			}
		})
	}
}
