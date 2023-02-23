package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "."
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		fmt.Println(path)
	}
	return nil
}
