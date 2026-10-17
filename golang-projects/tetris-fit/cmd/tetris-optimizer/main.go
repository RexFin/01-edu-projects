package main

import (
	"fmt"
	"log"
	"os"
	"tetris-fit/internal/analyzer"
	"tetris-fit/internal/parser"
	"tetris-fit/internal/solver"
	"tetris-fit/internal/tetromino"
	"time"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println()
		fmt.Println("error: not valid arguments!")
		fmt.Println("example: go run cmd/tetris-optimizer/main.go <sample.txt> [--analyze]")
		fmt.Println()
		return
	}

	filePath := os.Args[1]

	sliceOfShapes, err := parser.Parse(filePath)
	if err != nil {
		log.Fatal(err)
	}

	shapeList := make([][]tetromino.Shape, 0, len(sliceOfShapes))

	for _, v := range sliceOfShapes {
		shapeList = append(shapeList, tetromino.New(v))
	}

	startTime := time.Now()

	board, stats := solver.Solve(shapeList, nil)

	executionTime := time.Since(startTime)

	board.Print()

	if len(os.Args) == 3 && os.Args[2] == "--analyze" {
		// Сначало собираем данные
		baseReport := analyzer.Report{
			FileName:         os.Args[1],
			TetrominoesCount: len(shapeList),
			ExecutionTime:    executionTime,
			Stats:            stats,
		}

		// С собраных данных делаем вычисление
		analysisData := analyzer.CalculateAnalysis(*board, baseReport)

		// Печатаем
		analyzer.PrintReport(analysisData)
	}

}
