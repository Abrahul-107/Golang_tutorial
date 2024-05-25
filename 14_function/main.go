package main

import (
	"fmt"
)

func main() {
	fmt.Println("Function in golang")
	fmt.Println(add("Rahul", "parida"))
	fmt.Println(adderInloop(1, 2, 3, 5, 6, 6, 7))
	fmt.Println(nextInt([]byte{'3', '4', '5', '6'}, 0))
}

func add(first_num string, second_num string) string {
	return first_num + " " + second_num
}
func adderInloop(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total

}
func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func isDigit(b byte) bool {
	if b > 0 {
		return true
	} else {
		return false
	}
}
