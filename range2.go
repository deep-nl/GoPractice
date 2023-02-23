package main

import "fmt"

func main() {
	m := map[string]int{
		"rust": 1,
		"go":   2,
		"cpp":  3,
	}

	for k, v := range m {
		m[k+"_new"] = v + 10
	}

	// 每次运行的结果是不一样的，这是因为map是一个动态的数据
	fmt.Printf("map length: %d\n", len(m))

	// 如何解决这个问题，复制一份就可以
	m_copy := make(map[string]int, len(m))
	for k, v := range m {
		m_copy[k] = v
	}

	for k, v := range m_copy {
		m[k+"_newing"] = v + 10
	}
}
