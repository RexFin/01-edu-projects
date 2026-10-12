package analyzer

import (
	"math"
)

func CompactnessPercent(filledCells, totalCells float64) float64 {
	return math.Trunc(((filledCells/totalCells)*100)*100) / 100
}

func TheoreticalMinimum(tetrominoes int) int {
	return int(math.Ceil(math.Sqrt(float64(tetrominoes) * 4)))
}

func FilledCells(tetrominoes int) int {
	return tetrominoes * 4
}

func EmptyCells(boardSize, tetrominoes int) int {
	totalCells := boardSize * boardSize
	return totalCells - FilledCells(tetrominoes)
}
