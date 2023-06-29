package main

import (
	"fmt"
	"os"
)

func main() {
	// Retrieve the command-line arguments excluding the program name
	args := os.Args[1:]

	fmt.Println("CLI App - List of Strings")

	// Check if any arguments were provided
	if len(args) == 0 {
		fmt.Println("No strings provided.")
		return
	}

	fmt.Println("List of strings:")
	for _, arg := range args {
		fmt.Println(arg)
	}
}
