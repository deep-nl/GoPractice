package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("test reader's usage")
	buf := make([]byte, 4) //生成一个能够存放4 bytes数据的数组
	for {                  //无限循环直至有错或数据读取完返回EOF
		count, err := reader.Read(buf) //后面读取的内容会覆盖前面的buf的内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF : ", count)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(count, string(buf[:count]))
	}
}
