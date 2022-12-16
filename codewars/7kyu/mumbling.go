package main

import (
	"fmt"
	"strings"
)

func Accum(s string) (res string) {
	chars := strings.Split(s, "")

	for i, v := range chars {
		res += strings.ToUpper(v) + strings.ToLower(strings.Repeat(v, i)) + "-"
	}

	return res[:len(res)-1]
}

func main() {
	fmt.Println(Accum("ZpglnRxqenU") == "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu")
	fmt.Println(Accum("NyffsGeyylB") == "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb")
}
