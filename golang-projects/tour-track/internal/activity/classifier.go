// Package activity contains the deterministic, rule-based tour activity classifier.
package activity

// Score applies the project weights and caps the career longevity bonus at 10.
func Score(tourSpread, concertCount, careerYears int) int {
	if tourSpread < 0 {
		tourSpread = 0
	}
	if concertCount < 0 {
		concertCount = 0
	}
	bonus := careerYears / 5
	if bonus < 0 {
		bonus = 0
	}
	if bonus > 10 {
		bonus = 10
	}
	return tourSpread*2 + concertCount*3 + bonus
}

// Label maps a score to the thresholds specified by the auditor.
func Label(tourSpread, concertCount, careerYears int) string {
	return LabelForScore(Score(tourSpread, concertCount, careerYears))
}

func LabelForScore(score int) string {
	switch {
	case score >= 41:
		return "Legendary"
	case score >= 26:
		return "Established"
	case score >= 11:
		return "Active"
	default:
		return "Emerging"
	}
}

func Classify(tourSpread, concertCount, careerYears int) (int, string) {
	score := Score(tourSpread, concertCount, careerYears)
	return score, LabelForScore(score)
}

// IsRisingStar compares the latest two-year concert rate with the career average.
// The recent rate must exceed, not merely equal, 150% of the career average.
func IsRisingStar(recentConcertCount, concertCount, careerYears int) bool {
	if recentConcertCount <= 0 || concertCount < 0 || careerYears <= 0 {
		return false
	}
	recentRate := float64(recentConcertCount) / 2
	careerRate := float64(concertCount) / float64(careerYears)
	return recentRate > careerRate*1.5
}
