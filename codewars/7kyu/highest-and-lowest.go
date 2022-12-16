package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	numbers := strings.Fields(in)

	sort.Slice(numbers, func(i, j int) bool {
		num1, _ := strconv.Atoi(numbers[i])
		num2, _ := strconv.Atoi(numbers[j])
		return num1 > num2
	})

	return fmt.Sprintf("%s %s", numbers[0], numbers[len(numbers)-1])
}

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4") == "42 -9")
	fmt.Println(HighAndLow("1 2 3") == "3 1")
}
