package main

import (
	"fmt"
	"strings"
)

func CountBits(num uint) int {
	s := fmt.Sprintf("%b", num)
	count := strings.Count(s, "1")
	return int(count)
}

func main() {
	fmt.Println(CountBits(0))  // 0
	fmt.Println(CountBits(4))  // 1
	fmt.Println(CountBits(7))  // 3
	fmt.Println(CountBits(9))  // 2
	fmt.Println(CountBits(10)) // 2
}
