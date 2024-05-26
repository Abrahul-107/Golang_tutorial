package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	conversion()

}

// conversion function demonstrates how to read user input, convert it to a float64, and perform arithmetic operations.
func conversion() {

	// Create a new reader to read from standard input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your input")

	// Read input from the user until a newline character ('\n') is encountered.
	input, _ := reader.ReadString('\n')
	fmt.Println("Your input is ", input)

	// Convert the input string to a float64, trimming any leading or trailing whitespace.
	conversion, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// Check if there was an error during conversion.
	if err != nil {
		// Print the error if conversion failed.
		fmt.Println(err)
	} else {
		// Perform arithmetic operation and print the result.
		fmt.Println("Added 1 ", conversion+1)
	}
}
