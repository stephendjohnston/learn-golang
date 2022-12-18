package main

import (
	"fmt"
	"math"
)

func FindOutlier(integers []int) int {
	var even, odd, numOdd, numEven int
	for _, n := range integers {
		if n%2 == 0 {
			even++
			numEven = n
		} else {
			odd++
			numOdd = n
		}
	}

	if even > odd {
		return numOdd
	}

	return numEven
}

func main() {
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3}))                                                 // 3
	fmt.Println(FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781})) // 206847684
	fmt.Println(FindOutlier([]int{math.MaxInt32, 0, 1}))                                             // 0
}
