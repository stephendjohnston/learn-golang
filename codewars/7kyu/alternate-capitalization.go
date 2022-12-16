package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string) []string {
	s1 := ""
	s2 := ""

	for i, l := range st {
		letter := string(l)
		if i%2 == 0 {
			s1 += strings.ToUpper(letter)
			s2 += letter
		} else {
			s2 += strings.ToUpper(letter)
			s1 += letter
		}
	}

	return []string{s1, s2}
}

func main() {
	fmt.Println(Capitalize("abcdef"))       // []string{"AbCdEf", "aBcDeF"}
	fmt.Println(Capitalize("codewars"))     // []string{"CoDeWaRs", "cOdEwArS"}
	fmt.Println(Capitalize("abracadabra"))  // []string{"AbRaCaDaBrA", "aBrAcAdAbRa"}
	fmt.Println(Capitalize("codewarriors")) // []string{"CoDeWaRrIoRs", "cOdEwArRiOrS"}
}
