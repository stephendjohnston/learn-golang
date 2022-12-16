package main

import (
	"fmt"
)

// simple solution is just n * n * n
func RowSumOddNumbers(n int) int {
	m := make(map[int][]int)
	sliceLength := 1
	sliceVal := 1

	for i := 1; i <= n; i++ {
		for j := 0; j < sliceLength; j++ {
			m[i] = append(m[i], sliceVal)
			sliceVal += 2
		}
		sliceLength++
	}

	sum := 0

	for _, v := range m[n] {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(RowSumOddNumbers(1) == 1)
	fmt.Println(RowSumOddNumbers(2) == 8)
	fmt.Println(RowSumOddNumbers(13) == 2197)
	fmt.Println(RowSumOddNumbers(19) == 6859)
	fmt.Println(RowSumOddNumbers(41) == 68921)
	fmt.Println(RowSumOddNumbers(42) == 74088)
	fmt.Println(RowSumOddNumbers(74) == 405224)
}
