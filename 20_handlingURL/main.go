package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://leetcode.com/u/Rahul_parida/"

func main() {
	fmt.Println("Handling URL")
	fmt.Println(myurl)

	//Parsing
	result, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("No error")

	} else {
		fmt.Println(err)
	}
	fmt.Println("result sceme ", result.Scheme)
	fmt.Println("result host", result.Host)
	fmt.Println("result path", result.Path)
	fmt.Println("result port", result.Port())
	fmt.Println("result rawquery", result.RawQuery)

}
