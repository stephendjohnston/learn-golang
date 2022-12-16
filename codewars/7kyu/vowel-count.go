package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, v := range vowels {
		count += strings.Count(str, v)
	}

	return count
}

func main() {
	fmt.Println(GetCount("abracadabra") == 5)                         // true
	fmt.Println(GetCount("pear tree") == 4)                           // true
	fmt.Println(GetCount("o a kak ushakov lil vo kashu kakao") == 13) // true
}
