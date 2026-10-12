package parser

import (
	"fmt"
	"os"
	"strings"
)

func Parse(path string) ([][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", path, err)
	}

	// Удаляем мусор от windows (\r)
	index := 0
	for _, v := range data {
		if v != '\r' {
			data[index] = v
			index++
		}
	}
	data = data[:index]

	sliceOfRawShapes := strings.Split(string(data), "\n")
	sliceOfShapes := make([][]string, 0, len(sliceOfRawShapes)/5)
	currentShape := make([]string, 0, 4)

	blocks := 0
	for i, v := range sliceOfRawShapes {

		for j := 0; j < len(v); j++ {
			if v[j] == '#' {
				blocks++
			}

			if !(v[j] == '#' || v[j] == '.') {
				return nil, fmt.Errorf("ERROR")
			}
		}

		if (i+1)%5 == 0 {
			// Проверка является ли данная строка пустым (если там что-то другое то выводим ошибку)
			if v != "" {
				return nil, fmt.Errorf("ERROR")
			}

			if blocks != 4 {
				return nil, fmt.Errorf("ERROR")
			}

			// Если кусочки расположены криво, то это ошибка
			if len(currentShape) != 4 {
				return nil, fmt.Errorf("ERROR")
			}

			if !isBlockPosVaild(currentShape) {
				return nil, fmt.Errorf("ERROR")
			}

			sliceOfShapes = append(sliceOfShapes, currentShape)
			currentShape = []string{}
			blocks = 0
		} else {
			if len(v) != 4 {
				return nil, fmt.Errorf("ERROR")
			}
			currentShape = append(currentShape, v)
		}
	}

	// Последний элемент добавляем
	if len(currentShape) != 0 {
		if blocks != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		if len(currentShape) != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		if !isBlockPosVaild(currentShape) {
			return nil, fmt.Errorf("ERROR")
		}

		sliceOfShapes = append(sliceOfShapes, currentShape)
	}

	return sliceOfShapes, nil
}

func isBlockPosVaild(block []string) bool {
	connections := 0

	for i, s := range block {
		for j, v := range s {
			if v == '#' {
				if j+1 < 4 && s[j+1] == '#' {
					connections++
				}

				if i+1 < 4 && block[i+1][j] == '#' {
					connections++
				}
			}
		}
	}
	if connections == 4 || connections == 3 {
		return true
	} else {
		return false
	}
}
