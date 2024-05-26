package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map in golang")
	languages := make(map[string]int)

	languages["cpp"] = 107
	languages["py"] = 108
	languages["go"] = 109
	languages["java"] = 110
	languages["Dt"] = 111

	delete(languages, "Dt")
	fmt.Print(languages)

	for _, value := range languages {
		fmt.Printf("For key  , value is %v \n ", value)
	}
}
