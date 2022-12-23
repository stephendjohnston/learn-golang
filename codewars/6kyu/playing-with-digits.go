package main

import (
	"fmt"
	"math"
)

func DigPow(n, p int) int {
	N := n
	var digits []int

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	sum := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum += int(math.Pow(float64(digits[i]), float64(p)))
		p++
	}

	if sum%N == 0 {
		return sum / N
	} else {
		return -1
	}
}

func main() {
	fmt.Println(DigPow(89, 1))    // 1
	fmt.Println(DigPow(92, 1))    // -1
	fmt.Println(DigPow(695, 2))   // 2
	fmt.Println(DigPow(46288, 3)) // 51
}
