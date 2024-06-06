package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Get Request in GOLANG")
	performGetRequest()
}

func performGetRequest() {
	const myurl = "https://www.geeksforgeeks.org/amazon-sde-sheet-interview-questions-and-answers/?ref=ghm"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code is ", response.StatusCode)
	fmt.Println("Content length is ", response.ContentLength)

	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
}
