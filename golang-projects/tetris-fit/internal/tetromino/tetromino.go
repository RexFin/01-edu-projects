package tetromino

type Shape struct {
	X, Y int
}

func New(rawShape []string) []Shape {
	newCord := make([]Shape, 0, 4)

	for i := 0; i < len(rawShape); i++ {
		for j := 0; j < len(rawShape[i]); j++ {
			if rawShape[i][j] == '#' {
				newCord = append(newCord, Shape{j, i})
			}
		}
	}

	if len(newCord) == 0 {
		return newCord
	}

	minX := newCord[0].X
	minY := newCord[0].Y

	for i := 1; i < len(newCord); i++ {
		if newCord[i].X < minX {
			minX = newCord[i].X
		}
		if newCord[i].Y < minY {
			minY = newCord[i].Y
		}
	}

	for i := 0; i < len(newCord); i++ {
		newCord[i].X -= minX
		newCord[i].Y -= minY
	}

	return newCord
}
