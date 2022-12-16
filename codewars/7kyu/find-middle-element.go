package main

import (
	"fmt"
	"sort"
)

func Gimme(array [3]int) int {
	sliceCopy := make([]int, len(array))
	copy(sliceCopy, array[:])
	sort.Ints(sliceCopy)
	m := map[int]int{array[0]: 0, array[1]: 1, array[2]: 2}

	return m[sliceCopy[1]]
}

func main() {
	fmt.Println(Gimme([3]int{2, 3, 1}) == 0)
	fmt.Println(Gimme([3]int{5, 10, 14}) == 1)
	fmt.Println(Gimme([3]int{1, 3, 4}) == 1)
	fmt.Println(Gimme([3]int{15, 10, 14}) == 2)
	fmt.Println(Gimme([3]int{-4, -23, 4}) == 0)
	fmt.Println(Gimme([3]int{-15, -10, 14}) == 1)
}
