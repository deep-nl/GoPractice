package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func dealLarge() {
	// Open the compressed file for reading
	file, err := os.Open("file.gz")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a gzip reader from the file
	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer reader.Close()

	// Use a buffer to read the uncompressed data in chunks
	bufferSize := 1024
	buffer := make([]byte, bufferSize)
	var data []byte
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading uncompressed data:", err)
			return
		}
		data = append(data, buffer[:n]...)
	}
	fmt.Printf("Uncompressed data: %s\n", data[:])
}

func dealLarge2() {
	// Open the gzip file
	file, err := os.Open("example.gz")
	if err != nil {
		fmt.Println("Error opening gzip file:", err)
		return
	}
	defer file.Close()

	// Create a gzip reader
	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer reader.Close()

	// Create a buffer to store the uncompressed data
	var buffer bytes.Buffer

	// Read and append the uncompressed data to the buffer
	data := make([]byte, 1024)
	for {
		n, err := reader.Read(data)
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			fmt.Println("Error reading uncompressed data:", err)
			return
		}
		buffer.Write(data[:n])
	}

	// Print the uncompressed data
	fmt.Printf("Uncompressed data: %s\n", buffer.String())
}

func dealSmall() {
	filename := "myfile.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))
}

func dealSmall2() {
	filename := "myfile.txt"

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Read the file in chunks
	chunkSize := 1024
	data := make([]byte, chunkSize)
	for {
		n, err := reader.Read(data)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		if n == 0 {
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, data[:n])
	}
	fmt.Printf("Uncompressed data: %s\n", data[:])
}

// io.Copy is  a useful function in many scenarios,
// especially when dealing with large amounts of data that can be read in a single operation.
func dealSmall3() {
	filename := "myfile.txt"

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Read and print the uncompressed data
	var data bytes.Buffer
	_, err = io.Copy(&data, reader)
	if err != nil {
		fmt.Println("Error reading uncompressed data:", err)
		return
	}
	fmt.Printf("Uncompressed data: %s\n", data.Bytes())
}
