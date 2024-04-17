package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{"apple", "banana", "cherry"}
	fmt.Println("Before append ", fruitlist)
	fruitlist = append(fruitlist, "coconout", "mango")
	fmt.Println("After append ", fruitlist)
	fmt.Println(fruitlist[1:3])

	highscores := make([]int, 4)
	highscores[0] = 1
	highscores[1] = 5
	highscores[2] = 9
	highscores[3] = 90

	highscores = append(highscores, 56, 34, 4)
	sort.Ints(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	fmt.Println(highscores)

}
