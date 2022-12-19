package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) (res string) {
	lowerWord := strings.ToLower(word)
	for _, r := range lowerWord {
		if strings.Count(lowerWord, string(r)) > 1 {
			res += ")"
		} else {
			res += "("
		}
	}
	return
}

func main() {
	fmt.Println(DuplicateEncode("din") == "(((")
	fmt.Println(DuplicateEncode("recede") == "()()()")
	fmt.Println(DuplicateEncode("(( @") == "))((")
}
