package main

import "fmt"

func Rps(p1, p2 string) string {
	if (p1 == "scissors" && p2 == "paper") || (p1 == "rock" && p2 == "scissors") || (p1 == "paper" && p2 == "rock") {
		return "Player 1 won!"
	}

	if (p2 == "scissors" && p1 == "paper") || (p2 == "rock" && p1 == "scissors") || (p2 == "paper" && p1 == "rock") {
		return "Player 2 won!"
	}
	return "Draw!"
}

func main() {
	fmt.Println(Rps("paper", "rock")) // Player 1 won!
}
