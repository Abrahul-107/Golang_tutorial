package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	userInput()
}

func userInput() {
	// Create a new reader to read from standard input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your roll no ")

	// Read input from the user until a newline character ('\n') is encountered.
	// We ignore the second return value (the error) using "_" because we're not handling errors for simplicity.
	input, _ := reader.ReadString('\n')
	fmt.Print(input)

	// Print the type of the input.
	fmt.Printf("Type is  %T", input)
}
