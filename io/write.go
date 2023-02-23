package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var writer bytes.Buffer

	strings := []string{
		"hello ",
		"right now i am testing the usage of ",
		"writer",
	}
	for _, s := range strings {
		n, err := writer.Write([]byte(s))
		if err != nil {
			log.Fatal(err)
		}
		if n != len(s) {
			log.Fatal("fail to write the right string")
		}
	}
	fmt.Println(writer.String())
}