package solver

import (
	"fmt"
	"tetris-fit/internal/tetromino"
)

func SolveRecursive(index int, shapeList [][]tetromino.Shape, board *Board, stats *Stats, onStep StepCallback) bool {
	if index == len(shapeList) {
		return true
	}

	for y := 0; y < board.Size; y++ {
		for x := 0; x < board.Size; x++ {
			if board.CanPlace(shapeList[index], x, y) {
				board.Place(shapeList[index], x, y, 'A'+rune(index))
				stats.PlacementAttempts++

				if onStep != nil {
					msg := fmt.Sprintf("Step: place piece %c", 'A'+rune(index))
					onStep(board, stats, msg)
				}

				if SolveRecursive(index+1, shapeList, board, stats, onStep) == true {
					return true
				}

				board.Remove(shapeList[index], x, y)
				stats.Backtracks++

				if onStep != nil {
					msg := fmt.Sprintf("Backtrack: removed piece %c", 'A'+rune(index))
					onStep(board, stats, msg)
				}
			}
		}
	}

	return false
}
