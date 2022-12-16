package main

import (
	"fmt"
	"strings"
)

func ReverseWord(word string) string {
	chars := strings.Split(word, "")
	reversed := ""

	for _, char := range chars {
		reversed = string(char) + reversed
	}

	return reversed
}

func SpinWords(str string) string {
	words := strings.Fields(str)
	result := make([]string, len(words))

	for _, word := range words {
		if len(word) >= 5 {
			word = ReverseWord(word)
		}

		result = append(result, word)
	}

	return strings.Trim(strings.Join(result, " "), " ")
}

func main() {
	fmt.Println(SpinWords("Welcome") == "emocleW")
	fmt.Println(SpinWords("to") == "to")
	fmt.Println(SpinWords("CodeWars") == "sraWedoC")
	fmt.Println(SpinWords("Hey fellow warriors") == "Hey wollef sroirraw")
	fmt.Println(SpinWords("Burgers are my favorite fruit") == "sregruB are my etirovaf tiurf")
	fmt.Println(SpinWords("Pizza is the best vegetable") == "azziP is the best elbategev")
}
