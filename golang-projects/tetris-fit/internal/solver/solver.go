package solver

import (
	"fmt"
	"math"
	"tetris-fit/internal/tetromino"
)

type Board struct {
	Grid [][]rune
	Size int
}

type Stats struct {
	PlacementAttempts int
	Backtracks        int
}

// StepCallback это контракт. Любая функция с такими аргументами может быть визуализатором.
type StepCallback func(board *Board, stats *Stats, message string)

func MakeTetroBoard(size int) *Board {
	grid := make([][]rune, size)

	for i := range size {
		grid[i] = make([]rune, size)

		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	return &Board{
		Grid: grid,
		Size: size,
	}
}

func Solve(shapeList [][]tetromino.Shape, onStep StepCallback) (*Board, Stats) {
	// Для анализа
	stats := &Stats{
		PlacementAttempts: 0,
		Backtracks:        0,
	}

	isSolved := false
	size := int(math.Ceil(math.Sqrt(float64(len(shapeList) * 4))))

	for {
		board := MakeTetroBoard(size)
		isSolved = SolveRecursive(0, shapeList, board, stats, onStep)

		if isSolved {
			return board, *stats
		}

		size++
	}
}

func (b *Board) CanPlace(shape []tetromino.Shape, xPos, yPos int) bool {
	for _, s := range shape {
		if !((b.Size > xPos+s.X && b.Size > yPos+s.Y) && b.Grid[yPos+s.Y][xPos+s.X] == '.') {
			return false
		}
	}

	return true
}

func (b *Board) Place(shape []tetromino.Shape, xPos, yPos int, idSymbol rune) {
	for _, s := range shape {
		b.Grid[s.Y+yPos][s.X+xPos] = idSymbol
	}
}

func (b *Board) Remove(shape []tetromino.Shape, xPos, yPos int) {
	for _, s := range shape {
		b.Grid[s.Y+yPos][s.X+xPos] = '.'
	}
}

func (b *Board) Print() {
	for i := range b.Grid {
		for _, v := range b.Grid[i] {
			fmt.Print(string(v))
		}

		fmt.Println()
	}
}
