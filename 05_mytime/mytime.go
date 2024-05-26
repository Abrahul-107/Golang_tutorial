package main

import (
	"fmt"
	"time"
)

func main() {
	myTime()

}

func myTime() {

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// start := time.Now()
	// //... operation that takes 20 milliseconds ...
	// t := time.Now()
	// elapsed := t.Sub(start)
	// fmt.Println(elapsed)

	presentDateCreate := time.Date(2024, time.April, 7, 15, 15, 0, 0, time.UTC)
	fmt.Println(presentDateCreate)
	fmt.Println(presentDateCreate.Format("01-02-2006 Monday "))
}
