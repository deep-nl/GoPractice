package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("abcde")
	// 或者是 bytes.NewReader([]byte("abcde"))
	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Println(n, err, buf[:n])
		if err == io.EOF {
			break
		}
	}
}
