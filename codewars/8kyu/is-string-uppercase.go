package main

import (
	"fmt"
	"strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	return MyString(strings.ToUpper(string(s))) == s
}

func main() {
	fmt.Println(MyString("s").IsUpperCase())                          // false
	fmt.Println(MyString("ABCDEFGHIJKLMNOPQRSTUVWXYz").IsUpperCase()) // false
	fmt.Println(MyString("WHAT DOES THE FOX SAY").IsUpperCase())      // true
}
