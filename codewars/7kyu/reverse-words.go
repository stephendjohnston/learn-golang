package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	reversed := make([]string, len(words))

	for _, word := range words {
		reverseWord := ""
		for j := len(word) - 1; j >= 0; j-- {
			reverseWord += string(word[j])
		}
		reversed = append(reversed, reverseWord)
	}

	return strings.Trim(strings.Join(reversed, " "), " ") // reverse those words
}

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog.") == "ehT kciuq nworb xof spmuj revo eht yzal .god")
	fmt.Println(ReverseWords("apple") == "elppa")
	fmt.Println(ReverseWords("a b c d") == "a b c d")
	fmt.Println(ReverseWords("double  spaced  words") == "elbuod  decaps  sdrow")
	fmt.Println(ReverseWords("stressed desserts") == "desserts stressed")
	fmt.Println(ReverseWords("hello hello") == "olleh olleh")
}
