package main

import "fmt"

type some struct{}

func NewSome() interface{} {
	var c *some = nil
	return c
}

func main() {
	p := NewSome()
	if p == nil {
		fmt.Println("it is nil")
	} else {
		fmt.Println("it is not nil")
	}
	// 结论，nil 指针 != nil 接口，
	// interface 是一个16个字节的结构体，首8个字节是类型字段，后8个字节是数据指针。
	// nil 指针占用的内存大小取决于指针数据类型，通常等于代码运行的平台上的内存地址的大小。
	// 这个例子实际上是把空指针赋值向一个空接口，这样其前8个字节不等于空值
}
