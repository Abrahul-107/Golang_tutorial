package main

import (
	"fmt"
)

func main() {

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	/**

	Traditional loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	Loop with range
	for i := range days {
		fmt.Println(days[i])
	}

	Loop with index
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index, day)
	}

	Loop work like while in for

	values := 1

	for values < 10 {
		fmt.Println(values)
		values++
	}

	**/

	values := 1

	for values < 10 {
		fmt.Println(values)
		if values == 5 {
			break

		}
		if values == 2 {
			goto mylabel
		}
		values++
	}

mylabel:
	fmt.Println("Your are in the lable and your number is 2")

}
