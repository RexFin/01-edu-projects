package main

import (
	"fmt"
	"log"
	"os"
	"tetris-fit/internal/parser"
	"tetris-fit/internal/solver"
	"tetris-fit/internal/tetromino"
	"tetris-fit/internal/visualizer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
		fmt.Println("error: not valid arguments!")
		fmt.Println("example: go run cmd/tetris-visualizer/main.go <sample.txt>")
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

	board, _ := solver.Solve(shapeList, visualizer.AnimateStep)

	board.Print()

}
