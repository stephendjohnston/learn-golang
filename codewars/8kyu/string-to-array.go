package main

import (
	"fmt"
	"strings"
)

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}

func main() {
	fmt.Println(StringToArray("Robin Singh"))                        // [Robin Singh]
	fmt.Println(StringToArray("CodeWars"))                           // [CodeWars]
	fmt.Println(StringToArray("I love arrays they are my favorite")) // [I love arrays they are my favorite]
	fmt.Println(StringToArray("1 2 3"))                              // [1 2 3]
}
