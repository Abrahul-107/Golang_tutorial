package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrayBasic()
}
func arrayBasic() {

	//Array of static length given by programmer
	var arr = [3]int{1, 2, 3}
	arr1 := [3]int{4, 5, 7}
	fmt.Println(arr)
	fmt.Println(arr1)

	//two arrays (arr1 and arr2) with inferred lengths
	dynamicarray := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(dynamicarray)

	//Access Elements of an Array
	fmt.Println(reflect.TypeOf(dynamicarray))
	fmt.Println(dynamicarray[4])
	dynamicarray[1] = 45
	fmt.Println(dynamicarray)

	//If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
	//Tip: The default value for int is 0, and the default value for string is "".
	arr11 := [5]int{}             //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr11)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//Initializing array with specific value
	arr111 := [5]int{1: 10, 2: 40}
	fmt.Println(arr111)
}
