package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	words := regexp.MustCompile("[-_]").Split(s, -1)
	for i, word := range words[1:] {
		words[i+1] = strings.Title(word)
	}
	return strings.Join(words, "")
}

func main() {
	fmt.Println(ToCamelCase("The_Stealth_Warrior")) // "TheStealthWarrior")
	fmt.Println(ToCamelCase("the-Stealth-Warrior")) // "theStealthWarrior")
}
