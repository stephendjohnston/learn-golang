package main

import "fmt"

func GetGrade(a, b, c int) rune {
	switch av := (a + b + c) / 3; {
	case av >= 90:
		return 'A'
	case av >= 80:
		return 'B'
	case av >= 70:
		return 'C'
	case av >= 60:
		return 'D'
	default:
		return 'F'
	}
}

func main() {
	fmt.Println(GetGrade(95, 85, 94) == 'A') // true
}
