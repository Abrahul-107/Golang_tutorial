package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "This needs to go in file - for write this"
	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}
	io.WriteString(file, content)

	length, err := io.WriteString(file, content)
	checkNilError((err))
	fmt.Println("Length is ", length)
	defer file.Close()
	readFile("./myfile.txt")

}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError((err))
	fmt.Println("Text data inside the file is \n ", string(databyte))
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
