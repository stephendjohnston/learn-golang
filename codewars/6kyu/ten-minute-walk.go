package main

import (
	"fmt"
)

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	var ns, we int

	for _, r := range walk {
		switch r {
		case 'n':
			ns++
		case 's':
			ns--
		case 'w':
			we++
		case 'e':
			we--
		}
	}

	return ns == 0 && we == 0
}

func main() {
	fmt.Println(IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}) == true)
	fmt.Println(IsValidWalk([]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}) == false)
	fmt.Println(IsValidWalk([]rune{'w'}) == false)
	fmt.Println(IsValidWalk([]rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}) == false)
	fmt.Println(IsValidWalk([]rune{'e', 'e', 'e', 'e', 'w', 'w', 's', 's', 's', 's'}) == false)
	fmt.Println(IsValidWalk([]rune{'w', 's', 'n', 's', 'n', 'n', 'e', 'w', 's', 's'}) == false)
}
