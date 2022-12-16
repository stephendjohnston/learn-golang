package main

import (
	"fmt"
)

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

func main() {
	fmt.Println(RemoveChar("eloquent") == "loquen")
	fmt.Println(RemoveChar("country") == "ountr")
	fmt.Println(RemoveChar("person") == "erso")
	fmt.Println(RemoveChar("ooopsss") == "oopss")
}
