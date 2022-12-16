package main

import (
	"fmt"
)

func SmallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]

	for _, n := range numbers[1:] {
		if n < smallest {
			smallest = n
		}
	}
	return smallest
}

func main() {
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2}) == 2)
	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100}) == -345)

}
