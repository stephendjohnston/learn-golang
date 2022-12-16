package main

import (
	"fmt"
	"strings"
)

func FakeBin(x string) string {
	digits := strings.Split(x, "")
	binary := ""

	for i := 0; i < len(digits); i++ {
		if digits[i] >= "5" {
			binary += "1"
		} else {
			binary += "0"
		}
	}

	return binary
}

func main() {
	fmt.Println(FakeBin("45385593107843568") == "01011110001100111") // true
}
