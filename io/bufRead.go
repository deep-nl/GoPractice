package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读文件
func ReadDataFromTxt() {
	txtFile, err := os.OpenFile("55555.txt", os.O_RDONLY, 0777) // O_TRUNC 清空重写
	if err != nil {
		fmt.Println("WriteDataToTxt os.OpenFile() err:", err)
		return
	}
	defer txtFile.Close()

	bufReader := bufio.NewReader(txtFile)

	for {
		data, _, err := bufReader.ReadLine() // 读一行日志
		if err == io.EOF {                   // 如果列表读完了，退出
			fmt.Println("数据读完了~")
			break
		}

		fmt.Println("data: ", string(data))
	}
	return
}
