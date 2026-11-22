package activity

import "testing"

func TestScore(t *testing.T) {
	cases := []struct {
		name                          string
		spread, concerts, years, want int
	}{
		{"weighted inputs", 3, 4, 20, 22},
		{"career bonus is capped", 0, 0, 100, 10},
		{"negative years add no bonus", 1, 2, -5, 8},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.spread, tc.concerts, tc.years); got != tc.want {
				t.Fatalf("Score() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestLabelBoundaries(t *testing.T) {
	cases := []struct {
		score int
		want  string
	}{{0, "Emerging"}, {10, "Emerging"}, {11, "Active"}, {25, "Active"}, {26, "Established"}, {40, "Established"}, {41, "Legendary"}}
	for _, tc := range cases {
		if got := LabelForScore(tc.score); got != tc.want {
			t.Errorf("LabelForScore(%d) = %q, want %q", tc.score, got, tc.want)
		}
	}
}

func TestLabelFormulaExample(t *testing.T) {
	if got := Score(5, 8, 30); got != 40 {
		t.Fatalf("Score() = %d, want 40", got)
	}
	if got := Label(5, 8, 30); got != "Established" {
		t.Fatalf("Label() = %q, want Established", got)
	}
}

func TestIsRisingStarExample(t *testing.T) {
	if !IsRisingStar(14, 40, 10) {
		t.Fatal("7 recent concerts per year should exceed a career average of 4 by more than 50%")
	}
	if IsRisingStar(10, 40, 10) {
		t.Fatal("5 recent concerts per year should not exceed a career average of 4 by more than 50%")
	}
	if IsRisingStar(12, 40, 10) {
		t.Fatal("a recent rate equal to 150% of the career average should not qualify")
	}
}
