package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

/*
ioutil 相关的方法
readAll：内部方法，读取所有数据
ReadAll：外部方法，读取所有数据
ReadFile：读取文件所有内容
WriteFile：写入文件内容，如果文件存在会先清空原有内容
ReadDir：返回文件夹下文件列表
nopCloser：将io.Reader 类型包装成 io.ReadCloser 类型
Discard：用于丢弃数据
*/
func main() {}

func readByFile() {
	data, err := ioutil.ReadFile("./file/test.txt")
	if err != nil {
		log.Fatal("err:", err)
		return
	}
	fmt.Println("data", string(data))
}

func writeFile() {
	err := ioutil.WriteFile("./file/write_test.txt", []byte("hello world!"), 0644)
	if err != nil {
		panic(err)
		return
	}
}
