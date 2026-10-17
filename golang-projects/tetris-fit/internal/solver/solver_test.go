package solver_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"tetris-fit/internal/parser"
	"tetris-fit/internal/solver"
	"tetris-fit/internal/tetromino"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantBoard solver.Board
	}{
		{
			name:    "Single Square",
			content: "##..\n##..\n....\n....\n",
			wantBoard: solver.Board{
				Size: 2,
				Grid: [][]rune{
					{'A', 'A'},
					{'A', 'A'},
				},
			},
		},
		{
			name:    "Two Horizontal Bars",
			content: "####\n....\n....\n....\n\n####\n....\n....\n....\n",
			wantBoard: solver.Board{
				Size: 4,
				Grid: [][]rune{
					{'A', 'A', 'A', 'A'},
					{'B', 'B', 'B', 'B'},
					{'.', '.', '.', '.'},
					{'.', '.', '.', '.'},
				},
			},
		},
		{
			name:    "Five Squares",
			content: strings.Repeat("##..\n##..\n....\n....\n\n", 4) + "##..\n##..\n....\n....\n",
			wantBoard: solver.Board{
				Size: 6,
				Grid: [][]rune{
					{'A', 'A', 'B', 'B', 'C', 'C'},
					{'A', 'A', 'B', 'B', 'C', 'C'},
					{'D', 'D', 'E', 'E', '.', '.'},
					{'D', 'D', 'E', 'E', '.', '.'},
					{'.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.'},
				},
			},
		},
		{
			name:    "36 Squares",
			content: strings.Repeat("##..\n##..\n....\n....\n\n", 35) + "##..\n##..\n....\n....\n",
			wantBoard: solver.Board{
				Size: 12,
				Grid: [][]rune{
					{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F', 'F'},
					{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F', 'F'},

					{'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L'},
					{'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L'},

					{'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R'},
					{'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R'},

					{'S', 'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X'},
					{'S', 'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X'},
					// 24 (Y), 25 (Z), затем 26 (A)
					{'Y', 'Y', 'Z', 'Z', 'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D'},
					{'Y', 'Y', 'Z', 'Z', 'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D'},

					{'E', 'E', 'F', 'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J'},
					{'E', 'E', 'F', 'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J'},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			path := filepath.Join(tempDir, "test.txt")

			err := os.WriteFile(path, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}

			sliceOfShapes, err := parser.Parse(path)
			if err != nil {
				t.Fatalf("parser error: %v", err)
			}

			shapeList := make([][]tetromino.Shape, 0, len(sliceOfShapes))
			for _, v := range sliceOfShapes {
				shapeList = append(shapeList, tetromino.New(v))
			}

			board, _ := solver.Solve(shapeList, nil)

			if !reflect.DeepEqual(*board, tt.wantBoard) {
				t.Errorf("\nwanted board: %+v\nget board:  %+v", tt.wantBoard, *board)
			}
		})
	}
}
