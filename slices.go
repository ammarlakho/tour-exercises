package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var matrix [][]uint8
	for i := 0; i < dy; i++ {
		matrix = append(matrix, make([]uint8, dx))
	}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = uint8((i * j) / 2)
		}
	}
	return matrix
}

func main() {
	pic.Show(Pic)
}
