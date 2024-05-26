package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer in golang")
	defer third()
	defer second()

	fmt.Println("First executed ")
}

func third() {
	fmt.Println("Third function")
}
func second() {
	fmt.Println("Second function")
}
