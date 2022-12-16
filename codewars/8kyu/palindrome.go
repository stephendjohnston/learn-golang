package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsPalindrome("a"))
	fmt.Println(IsPalindrome("aba"))
	fmt.Println(IsPalindrome("Abba"))
	fmt.Println(IsPalindrome("hello"))
}

func IsPalindrome(str string) bool {
	var result string
	for _, v := range str {
		result = string(v) + result
	}

	if strings.ToLower(result) == strings.ToLower(str) {
		return true
	}
	return false
}
