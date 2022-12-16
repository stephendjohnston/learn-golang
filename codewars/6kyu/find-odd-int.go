package main

import (
	"fmt"
)

func FindOdd(seq []int) int {
	counts := make(map[int]int, len(seq))

	for _, v := range seq {
		counts[v] += 1
	}

	for k, v := range counts {
		if v%2 == 1 {
			return k
		}
	}

	return -1
}

func main() {
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5})) // 5
	fmt.Println(FindOdd([]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}))                      // -1
	fmt.Println(FindOdd([]int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}))                 // 5
	fmt.Println(FindOdd([]int{10}))                                                      // 10
	fmt.Println(FindOdd([]int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}))                        // 10
	fmt.Println(FindOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}))                       // 1
}
