package main

import (
	"fmt"
)

func GetSum(a, b int) (sum int) {
	if a > b {
		a, b = b, a
	}

	for i := a; i <= b; i++ {
		sum += i
	}

	return
}

func main() {
	fmt.Println(GetSum(0, 1) == 1)
	fmt.Println(GetSum(1, 2) == 3)
	fmt.Println(GetSum(5, -1) == 14)
	fmt.Println(GetSum(505, 4) == 127759)
	fmt.Println(GetSum(321, 123) == 44178)
	fmt.Println(GetSum(0, -1) == -1)
	fmt.Println(GetSum(-50, 0) == -1275)
	fmt.Println(GetSum(-1, -5) == -15)
	fmt.Println(GetSum(-5, -5) == -5)
}
