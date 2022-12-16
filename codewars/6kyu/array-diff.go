package main

import (
	"fmt"
)

func ArrayDiff(a, b []int) (res []int) {
	mp := map[int]bool{}

	for _, v := range b {
		mp[v] = true
	}

	for _, v := range a {
		if !mp[v] {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	emptyArr := make([]int, 0)
	fmt.Println(ArrayDiff([]int{1, 2}, []int{1}))       //([]int{2})
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{1}))    //([]int{2,2})
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{2}))    //([]int{1})
	fmt.Println(ArrayDiff([]int{1, 2, 2}, emptyArr))    //([]int{1,2,2})
	fmt.Println(ArrayDiff(emptyArr, []int{1, 2}))       // BeEmpty()
	fmt.Println(ArrayDiff([]int{1, 2, 3}, []int{1, 2})) //([]int{3})
}
