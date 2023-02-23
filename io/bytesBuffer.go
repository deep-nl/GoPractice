package main

import (
	"bytes"
	"fmt"
	"os"
)

// bytes.Buffer 是一个结构体类型，用来暂存写入的数据，这个结构体实现了 io.Writer 接口的 Write 方法
func main() {
	// 创建 Buffer 暂存空间，并将一个字符串写入 Buffer
	// 使用 io.Writer 的 Write 方法写入
	var buf bytes.Buffer
	buf.Write([]byte("hello world , "))

	// 用 Fprintf 将一个字符串拼接到 Buffer 里
	fmt.Fprintf(&buf, " welcome to golang !")

	// 将 Buffer 的内容输出到标准输出设备
	buf.WriteTo(os.Stdout)
}
