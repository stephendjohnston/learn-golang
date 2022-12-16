package main

import (
	"fmt"
)

func Between(a, b int) (bt []int) {
	for ; a <= b; a++ {
		bt = append(bt, a)
	}

	return
}

func main() {
	fmt.Println(Between(1, 4))  //  == [1 2 3 4]
	fmt.Println(Between(-2, 2)) //  == [-2 -1 0 1 2]
}
