package main

import (
	"fmt"
	"strings"
)

func convertToString(numbers interface{}) string {
	fmt.Println(strings.Replace(fmt.Sprint(numbers), " ", "", -1))
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]"))
	return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

func CreatePhoneNumber(numbers [10]uint) string {
	str := convertToString(numbers)
	return fmt.Sprintf("(%s) %s-%s", str[0:3], str[3:6], str[6:10])
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})) //"(123) 456-7890"))
}
