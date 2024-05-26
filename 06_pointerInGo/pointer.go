package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println("Pointer in GOLANG")
	NumCpu()
	pointer()
}

func pointer() {
	// Declare a variable 'a' and assign it the value 6
	a := 6

	// Declare a pointer variable 'ptr' of type int and assign it the memory address of 'a'
	var ptr *int = &a

	// Print the memory address stored in 'ptr'
	fmt.Println(ptr)

	// Print the value stored at the memory address pointed to by 'ptr'
	fmt.Println(*ptr)

	// Print the value of 'a'
	fmt.Println(a)

	// Declare a new pointer variable 'newptr' and assign it the memory address of 'a'
	var newptr = &a

	// Dereference 'newptr' and increment the value stored at the memory address by 4
	*newptr += 4

	// Print the updated value of 'a'
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(newptr))
}
func NumCpu() {
	/*func runtime.NumCPU() int
	NumCPU returns the number of logical CPUs usable by the current process.*/
	fmt.Println("Number of CPUs:", runtime.NumCPU())
}
