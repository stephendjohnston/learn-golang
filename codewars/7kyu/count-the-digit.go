package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NbDig(n int, d int) (count int) {
	for i := 0; i <= n; i++ {
		count += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}

	return count
}

func main() {
	fmt.Println(NbDig(10, 1) == 4)
	fmt.Println(NbDig(25, 1) == 11)
	fmt.Println(NbDig(550, 5) == 213)
	fmt.Println(NbDig(5750, 0) == 4700)
}
