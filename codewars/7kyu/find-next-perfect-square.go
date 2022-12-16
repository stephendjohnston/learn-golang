package main

import (
	"fmt"
	"math"
)

func FindNextSquare(sq int64) int64 {
	square := math.Sqrt(float64(sq))

	if square != math.Trunc(square) {
		return -1
	}

	return int64(math.Pow(square+1, 2))
}

func main() {
	fmt.Println(FindNextSquare(int64(121)))
	fmt.Println(FindNextSquare(int64(625)))
	fmt.Println(FindNextSquare(int64(319225)))
	fmt.Println(FindNextSquare(int64(15241383936)))
	fmt.Println(FindNextSquare(int64(155)))
}
