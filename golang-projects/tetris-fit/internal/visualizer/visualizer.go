package visualizer

import (
	"fmt"
	"tetris-fit/internal/solver"
	"time"
)

func AnimateStep(board *solver.Board, stats *solver.Stats, message string) {
	// Очищаем консоль (магия ANSI-символов)
	fmt.Print("\033[H\033[2J")

	fmt.Println(message)

	board.Print()

	fmt.Printf("attempts: %d, backtracks: %d\n", stats.PlacementAttempts, stats.Backtracks)

	// Делаем паузу, чтобы глаз успел заметить кадр
	time.Sleep(200 * time.Millisecond)
}
