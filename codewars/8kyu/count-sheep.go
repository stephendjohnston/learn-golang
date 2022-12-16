package main

import (
	"fmt"
)

func countSheep(num int) string {
	sheep := ""

	for i := 1; i <= num; i++ {
		sheep += fmt.Sprintf("%v sheep...", i)
	}

	return sheep
}

func main() {
	fmt.Println(countSheep(2) == "1 sheep...2 sheep...")
}
