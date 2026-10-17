package analyzer

import (
	"fmt"
	"tetris-fit/internal/solver"
	"time"
)

type Report struct {
	FileName         string
	TetrominoesCount int
	ExecutionTime    time.Duration
	Stats            solver.Stats
}

// Итоговые данные для теста
type AnalysisData struct {
	FileName          string
	TetrominoesCount  int
	GridSize          int
	TotalCells        int
	FilledCells       int
	EmptyCells        int
	Compactness       float64
	TheoryMin         int
	MatchesTheoryMin  bool
	ExecutionTimeMs   int64
	PlacementAttempts int
	Backtracks        int
}

func CalculateAnalysis(board solver.Board, report Report) AnalysisData {
	filled := FilledCells(report.TetrominoesCount)
	total := board.Size * board.Size
	theoryMin := TheoreticalMinimum(report.TetrominoesCount)

	return AnalysisData{
		FileName:          report.FileName,
		TetrominoesCount:  report.TetrominoesCount,
		GridSize:          board.Size,
		TotalCells:        total,
		FilledCells:       filled,
		EmptyCells:        EmptyCells(board.Size, report.TetrominoesCount),
		Compactness:       CompactnessPercent(float64(filled), float64(total)),
		TheoryMin:         theoryMin,
		MatchesTheoryMin:  board.Size == theoryMin,
		ExecutionTimeMs:   report.ExecutionTime.Milliseconds(),
		PlacementAttempts: report.Stats.PlacementAttempts,
		Backtracks:        report.Stats.Backtracks,
	}
}

// Печатает только вывод
func PrintReport(data AnalysisData) {
	matchStatus := "differ"
	if data.MatchesTheoryMin {
		matchStatus = "matches"
	}

	fmt.Println("\n=== Solution Analysis ===")
	fmt.Println()
	fmt.Printf("Input file: %s\n", data.FileName)
	fmt.Printf("Tetrominoes: %d\n\n", data.TetrominoesCount)

	fmt.Println("Solution quality:")
	fmt.Printf("  Grid size: %dx%d (%d cells)\n", data.GridSize, data.GridSize, data.TotalCells)
	fmt.Printf("  Filled cells: %d\n", data.FilledCells)
	fmt.Printf("  Empty cells: %d\n", data.EmptyCells)
	fmt.Printf("  Compactness: %.2f%%\n", data.Compactness)
	fmt.Printf("  Theoretical minimum: %dx%d (%s)\n", data.TheoryMin, data.TheoryMin, matchStatus)

	fmt.Println("\nPerformance:")
	fmt.Printf("  Execution time: %dms\n", data.ExecutionTimeMs)
	fmt.Printf("  Placement attempts: %d\n", data.PlacementAttempts)
	fmt.Printf("  Backtracks: %d\n", data.Backtracks)
}
