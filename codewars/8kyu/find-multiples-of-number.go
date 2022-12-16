package main

import (
	"fmt"
)

func FindMultiples(integer, limit int) []int {
	s := make([]int, 0)

	for i := integer; i <= limit; i += integer {
		s = append(s, i)
	}

	return s
}

func main() {
	fmt.Println(FindMultiples(5, 25)) // [5 10 15 20 25]
}
