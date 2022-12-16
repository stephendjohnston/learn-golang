package main

import (
	"fmt"
)

func solution(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str)-len(ending):] == ending
}

func main() {
	fmt.Println(solution("", ""))              // true
	fmt.Println(solution(" ", ""))             // true
	fmt.Println(solution("abc", "c"))          // true
	fmt.Println(solution("banana", "ana"))     // true
	fmt.Println(solution("a", "z"))            // false
	fmt.Println(solution("", "t"))             // false
	fmt.Println(solution("$a = $b + 1", "+1")) // false
	fmt.Println(solution("    ", "   "))       // true
	fmt.Println(solution("1oo", "100"))        // false
}
