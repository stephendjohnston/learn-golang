package main

import "fmt"

func vals() (int, int) { // (int, int) in this function signature shows that the function returns 2 ints
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a) // 3
	fmt.Println(b) // 7

	_, c := vals()
	fmt.Println(c) // 7
}
