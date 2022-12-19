package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Persistence(n int) (mp int) {
	for n > 9 {
		mp++
		sl := strings.Split(strconv.Itoa(n), "")
		sum := 1
		for _, v := range sl {
			i, _ := strconv.Atoi(v)
			sum *= i
		}
		n = sum
		sum = 0
	}

	return mp
}

func main() {
	fmt.Println((Persistence(39)) == 3)
	fmt.Println((Persistence(4)) == 0)
	fmt.Println((Persistence(25)) == 2)
	fmt.Println((Persistence(999)) == 4)
}
