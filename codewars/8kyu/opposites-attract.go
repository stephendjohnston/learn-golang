package main

import (
	"fmt"
)

func LoveFunc(flower1, flower2 int) bool {
	return (flower1+flower2)%2 == 1
}

func main() {
	fmt.Println(LoveFunc(1, 4)) // true
	fmt.Println(LoveFunc(2, 2)) // false
}
