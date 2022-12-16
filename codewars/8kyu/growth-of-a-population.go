package main

import (
	"fmt"
)

func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0
	percent = percent / 100
	for p > p0 {
		p0 = int(float64(p0)*percent) + p0 + aug
		years++
	}

	return years
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000) == 15)             // true
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000) == 10)   // true
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000) == 94)   // true
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000) == 151) // true
}
