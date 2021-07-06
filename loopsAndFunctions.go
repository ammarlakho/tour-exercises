package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	g := 1.0
	sq := math.Sqrt(x)
	for math.Abs(g-sq) > 1e-10 {
		g -= (g*g - x) / (2 * g)
	}
	return g
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
