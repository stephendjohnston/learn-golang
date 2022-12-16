package main

import (
	"fmt"
)

func GetMiddle(s string) string {
	even := len(s)%2 == 0
	middle := len(s) / 2

	if even {
		return s[middle-1 : middle+1]
	}

	return s[middle : middle+1]
}

func main() {
	fmt.Println(GetMiddle("test") == "es")
	fmt.Println(GetMiddle("testing") == "t")
	fmt.Println(GetMiddle("middle") == "dd")
	fmt.Println(GetMiddle("A") == "A")
}
