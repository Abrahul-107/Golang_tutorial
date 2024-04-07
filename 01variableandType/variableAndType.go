package main

import (
	"fmt"
	"reflect"
)

func main() {
	variableType()
	variableDeclarationWithoutInitialization()
	multipleVariableDeclaration()

}

func variableType() {

	// You always have to specify either type or value (or both)
	var rahul string = "GoodBoy"
	fmt.Print(rahul, "\n")

	/*It is possible to assign a value to a variable after it is declared.
	This is helpful for cases the value is not initially known.*/

	var student1 string
	student1 = "John"
	fmt.Println(student1)

	/*Note: In this case, the type of the variable is inferred from the value
	(means that the compiler decides the type of the variable, based on the value).*/
	rahull := 107
	fmt.Println(rahull)

	var smallval uint8 = 255
	fmt.Println(smallval)

}

func variableDeclarationWithoutInitialization() {

	/*In Go, all variables are initialized. So, if you declare a variable without an initial value,
	its value will be set to the default value of its type:*/

	var a string
	var b int
	var c float32
	var d bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

func multipleVariableDeclaration() {

	//In Go, it is possible to declare multiple variables in the same line.
	a, b, c, d := 1, 2, 8, 5
	fmt.Println(a, b, c, d)
	fmt.Println(reflect.TypeOf(a))

	//Go Variable Declaration in a Block
	var (
		first  int     = 66
		second string  = "Rahul"
		third  float64 = 56.6
	)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(reflect.TypeOf(first))
}
