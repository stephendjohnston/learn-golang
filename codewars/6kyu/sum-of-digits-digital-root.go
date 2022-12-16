package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DigitalRoot(n int) int {
	for n >= 10 {
		digits := strings.Split(strconv.Itoa(n), "")
		sum := 0

		for _, strDig := range digits {
			intDig, _ := strconv.Atoi(strDig)
			sum += intDig
		}
		n = sum
	}

	return n
}

func main() {
	fmt.Println(DigitalRoot(16))     // 7
	fmt.Println(DigitalRoot(195))    // 6
	fmt.Println(DigitalRoot(992))    // 2
	fmt.Println(DigitalRoot(167346)) // 9
	fmt.Println(DigitalRoot(0))      // 0
}
