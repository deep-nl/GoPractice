package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
bufio.Scanner is a package in Go that provides a convenient way to read data from a file or input stream.
It reads data in the form of lines or words, and can be used to parse files and other text-based data.
*/

func main() {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	// Scan through each line of the file
	for scanner.Scan() {
		// Get the current line from the scanner
		line := scanner.Text()

		// Print the line
		fmt.Println(line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}
