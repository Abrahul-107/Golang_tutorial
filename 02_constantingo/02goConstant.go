package main

import (
	"fmt"
)

const rahul = "RAHUL"

func main() {
	constantInGo()
}

// Define constantInGo within the same package
func constantInGo() {
	const A = 10
	fmt.Println(A)
	fmt.Print(rahul)
}
