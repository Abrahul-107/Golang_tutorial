package main

import (
	"fmt"
)

func main() {

	second_num := 15

	if first_num := 156; first_num > second_num {
		fmt.Println(first_num, "is greater than ", second_num)
	} else {
		fmt.Println(second_num, "is greater than ", first_num)
	}
}
