package analyzer

import (
	"testing"
	"tetris-fit/internal/solver"
	"time"
)

func TestCalculateAnalysis(t *testing.T) {
	tests := []struct {
		name   string
		board  solver.Board
		report Report
		want   AnalysisData
	}{
		{
			name:  "Ideal 4x4 board",
			board: solver.Board{Size: 4},
			report: Report{
				FileName:         "testdata/goodexample02.txt",
				TetrominoesCount: 4,
				ExecutionTime:    5 * time.Millisecond,
				Stats: solver.Stats{
					PlacementAttempts: 42,
					Backtracks:        10,
				},
			},
			want: AnalysisData{
				FileName:          "testdata/goodexample02.txt",
				TetrominoesCount:  4,
				GridSize:          4,
				TotalCells:        16,
				FilledCells:       16,
				EmptyCells:        0,
				Compactness:       100.00,
				TheoryMin:         4,
				MatchesTheoryMin:  true,
				ExecutionTimeMs:   5,
				PlacementAttempts: 42,
				Backtracks:        10,
			},
		},
		{
			name:  "Expanded 5x5 board (differs from theoretical minimum)",
			board: solver.Board{Size: 5},
			report: Report{
				FileName:         "testdata/goodexample01.txt",
				TetrominoesCount: 4,
				ExecutionTime:    12 * time.Millisecond,
				Stats: solver.Stats{
					PlacementAttempts: 150,
					Backtracks:        45,
				},
			},
			want: AnalysisData{
				FileName:          "testdata/goodexample01.txt",
				TetrominoesCount:  4,
				GridSize:          5,
				TotalCells:        25,
				FilledCells:       16,
				EmptyCells:        9,
				Compactness:       64.00,
				TheoryMin:         4,
				MatchesTheoryMin:  false,
				ExecutionTimeMs:   12,
				PlacementAttempts: 150,
				Backtracks:        45,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAnalysis(tt.board, tt.report)

			if got.TotalCells != tt.want.TotalCells {
				t.Errorf("\nwanted TotalCells: %v\nget TotalCells:  %v", tt.want.TotalCells, got.TotalCells)
			}

			if got.Compactness != tt.want.Compactness {
				t.Errorf("\nwanted Compactness: %v\nget Compactness:  %v", tt.want.Compactness, got.Compactness)
			}

			if got.MatchesTheoryMin != tt.want.MatchesTheoryMin {
				t.Errorf("\nwanted MatchesTheoryMin: %v\nget MatchesTheoryMin:  %v", tt.want.MatchesTheoryMin, got.MatchesTheoryMin)
			}

			if got.FilledCells != tt.want.FilledCells {
				t.Errorf("\nwanted FilledCells: %v\nget FilledCells:  %v", tt.want.FilledCells, got.FilledCells)
			}

			if got.EmptyCells != tt.want.EmptyCells {
				t.Errorf("\nwanted EmptyCells: %v\nget EmptyCells:  %v", tt.want.EmptyCells, got.EmptyCells)
			}

			if got.ExecutionTimeMs != tt.want.ExecutionTimeMs {
				t.Errorf("\nwanted ExecutionTimeMs: %v\nget ExecutionTimeMs:  %v", tt.want.ExecutionTimeMs, got.ExecutionTimeMs)
			}
		})
	}
}
