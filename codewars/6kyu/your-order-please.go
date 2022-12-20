package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Order(sentence string) string {
	sl := strings.Fields(sentence)
	re := regexp.MustCompile(`[a-zA-Z]`)
	sort.Slice(sl, func(i, j int) bool {
		n1 := re.ReplaceAll([]byte(sl[i]), []byte(""))[0]
		n2 := re.ReplaceAll([]byte(sl[j]), []byte(""))[0]
		return n1 < n2
	})

	return strings.Join(sl, " ")
}

func main() {
	fmt.Println(Order("is2 Thi1s T4est 3a"))               //== "Thi1s is2 3a T4est"
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2")) //== "Fo1r the2 g3ood 4of th5e pe6ople"
	fmt.Println(Order(""))                                 //== ""
}
