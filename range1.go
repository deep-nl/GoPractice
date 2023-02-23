package main

import "fmt"

func main() {
	s := []string{"hello"}

	// 思考的这个循环为什么不是无限的
	for range s {
		s = append(s, "here")
		fmt.Print(s)
	}

	// 这样就是无限循环
	for i := 0; i < len(s); i++ {
		s = append(s, "there")
		fmt.Println(s)
	}
}
