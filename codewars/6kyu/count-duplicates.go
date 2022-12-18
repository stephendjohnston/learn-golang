package main

import (
	"fmt"
	"strings"
)

func duplicateCount(s1 string) int {
	mp := map[rune]int{}
	var count int

	for _, r := range strings.ToLower(s1) {
		mp[r]++
		if mp[r] == 2 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(duplicateCount("abcde"))          // 0
	fmt.Println(duplicateCount("abcdea"))         // 1
	fmt.Println(duplicateCount("abcdeaB11"))      // 3
	fmt.Println(duplicateCount("indivisibility")) // 1
}
