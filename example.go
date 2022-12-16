package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func main() {
	myStack := Stack{}
	yourStack := Stack{}
	fmt.Println(myStack, yourStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	yourStack.Push(100)
	yourStack.Push(200)
	yourStack.Push(300)
	fmt.Println(myStack, yourStack)
}
