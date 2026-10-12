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

func PrintReport(board solver.Board, report Report) {
	filled := FilledCells(report.TetrominoesCount)
	empty := EmptyCells(board.Size, report.TetrominoesCount)
	total := board.Size * board.Size
	compactness := CompactnessPercent(float64(filled), float64(total))
	theoryMin := TheoreticalMinimum(report.TetrominoesCount)

	fmt.Println("\n=== Solution Analysis ===")
	fmt.Println()

	fmt.Printf("Input file: %s\n", report.FileName)
	fmt.Printf("Tetrominoes: %d\n\n", report.TetrominoesCount)

	fmt.Println("Solution quality:")
	fmt.Printf("  Grid size: %dx%d (%d cells)\n", board.Size, board.Size, total)
	fmt.Printf("  Filled cells: %d\n", filled)
	fmt.Printf("  Empty cells: %d\n", empty)
	fmt.Printf("  Compactness: %.2f%%\n", compactness)

	// Проверка на совподение теоритической минимума с текущим размером доски
	if board.Size == theoryMin {
		fmt.Printf("  Theoretical minimum: %dx%d (matches)\n", theoryMin, theoryMin)
	} else {
		fmt.Printf("  Theoretical minimum: %dx%d (differ)\n", theoryMin, theoryMin)
	}

	fmt.Println("\nPerformance:")
	fmt.Printf("  Execution time: %vms\n", report.ExecutionTime.Milliseconds())
	fmt.Printf("  Placement attempts: %d\n", report.Stats.PlacementAttempts)
	fmt.Printf("  Backtracks: %d\n", report.Stats.Backtracks)
}
