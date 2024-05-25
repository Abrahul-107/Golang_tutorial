package main

import (
	"fmt"
)

type User struct {
	Name       string
	Email      string
	Age        int
	Attendence bool
}

func main() {
	fmt.Println("Structs in go")

	Rahul := User{"Rahul", "abrahul882@gmail.com", 23, true}
	fmt.Print(Rahul, "\n")
	fmt.Printf("%+v", Rahul)
}
