package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://codeforces.com"

func main() {
	fmt.Println("Web request")
	resonse, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response is", resonse)
	defer resonse.Body.Close()

	databytes, err := ioutil.ReadAll(resonse.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Whole content", string(databytes))
}
