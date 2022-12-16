package main

import (
	"fmt"
	"math"
)

func Movie(card, ticket int, perc float64) (n int) {
	totalA := ticket
	discountTickPrice := float64(ticket) * perc
	totalB := float64(card) + discountTickPrice
	n++

	for float64(totalA) <= math.Ceil(totalB) {
		totalA += ticket
		discountTickPrice = discountTickPrice * perc
		totalB += discountTickPrice
		n++
	}

	return
}

func main() {
	fmt.Println(Movie(500, 15, 0.9))  // 43
	fmt.Println(Movie(100, 10, 0.95)) // 24
	fmt.Println(Movie(0, 10, 0.95))   // 2
	fmt.Println(Movie(994, 2, 0.44))  // 499
}
