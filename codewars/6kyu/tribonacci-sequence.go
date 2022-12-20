package main

import (
	"fmt"
)

func Tribonacci(signature [3]float64, n int) (result []float64) {
	result = signature[:]

	for i := 0; i < n-3; i++ {
		next := result[i] + result[i+1] + result[i+2]
		result = append(result, next)
	}

	return result[:n]
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))       // []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105})
	fmt.Println(Tribonacci([3]float64{0, 0, 1}, 10))       // []float64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44})
	fmt.Println(Tribonacci([3]float64{0, 1, 1}, 10))       // []float64{0, 1, 1, 2, 4, 7, 13, 24, 44, 81})
	fmt.Println(Tribonacci([3]float64{1, 0, 0}, 10))       // []float64{1, 0, 0, 1, 1, 2, 4, 7, 13, 24})
	fmt.Println(Tribonacci([3]float64{0, 0, 0}, 10))       // []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(Tribonacci([3]float64{1, 2, 3}, 10))       // []float64{1, 2, 3, 6, 11, 20, 37, 68, 125, 230})
	fmt.Println(Tribonacci([3]float64{3, 2, 1}, 10))       // []float64{3, 2, 1, 6, 9, 16, 31, 56, 103, 190})
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 1))        // []float64{1})
	fmt.Println(Tribonacci([3]float64{300, 200, 100}, 0))  // []float64{})
	fmt.Println(Tribonacci([3]float64{0.5, 0.5, 0.5}, 30)) // []float64{0.5, 0.5, 0.5, 1.5, 2.5, 4.5, 8.5, 15.5, 28.5, 52.5, 96.5, 177.5, 326.5, 600.5, 1104.5, 2031.5, 3736.5, 6872.5, 12640.5, 23249.5, 42762.5, 78652.5, 144664.5, 266079.5, 489396.5, 900140.5, 1655616.5, 3045153.5, 5600910.5, 10301680.5})
}
