package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
	words := strings.Fields(str)

	for i, v := range words {
		words[i] = strings.ToUpper(string((v[0]))) + string(v[1:])
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ToJadenCase("most trees are blue"))
	fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
	fmt.Println(ToJadenCase("When I die. then you will realize"))
	fmt.Println(ToJadenCase("Jonah Hill is a genius"))
	fmt.Println(ToJadenCase("Dying is mainstream"))
}
