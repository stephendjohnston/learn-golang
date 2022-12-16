package main

import (
	"fmt"
)

func SquareSum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v * v
	}

	return
}

func main() {
	fmt.Println(SquareSum([]int{1, 2}) == 5)
	fmt.Println(SquareSum([]int{0, 3, 4, 5}) == 50)
}
