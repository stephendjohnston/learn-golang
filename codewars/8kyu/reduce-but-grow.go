package main

import "fmt"

func Grow(arr []int) (sumProduct int) {
	sumProduct = arr[0]
	for i := 1; i < len(arr); i++ {
		sumProduct *= arr[i]
	}
	return
}

func main() {
	fmt.Println(Grow([]int{1, 2, 3}))          // 6
	fmt.Println(Grow([]int{4, 1, 1, 1, 4}))    // 16
	fmt.Println(Grow([]int{2, 2, 2, 2, 2, 2})) // 64
}
