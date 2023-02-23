package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
)

// WriteDataToTxt bufio库的流式处理	如果文件较小，使用ioutil也不失为一种方法
func WriteDataToTxt() {
	txtFile, err := os.OpenFile("55555.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777) // O_TRUNC 清空重写
	if err != nil {
		fmt.Println("WriteDataToTxt os.OpenFile() err:", err)
		return
	}
	defer txtFile.Close()
	// txtFile.WriteString() // os操作文件-效率低

	bufWriter := bufio.NewWriter(txtFile)
	var wg sync.WaitGroup
	limitChan := make(chan struct{}, runtime.GOMAXPROCS(runtime.NumCPU())) // 最大并发协程数
	var mutex sync.Mutex

	for i := 0; i < 10000; i++ { // 写1w行测试
		limitChan <- struct{}{}
		wg.Add(1)

		go func(j int) {
			defer func() {
				if e := recover(); e != nil {
					fmt.Printf("WriteDataToTxt panic: %v,stack: %s\n", e, debug.Stack())
					// return
				}

				wg.Done()
				<-limitChan
			}()

			// 模拟业务逻辑：先整合所有数据，然后再统一写WriteString()
			strId := fmt.Sprintf("%v", j)
			strName := fmt.Sprintf(" user_%v", j)
			strScore := fmt.Sprintf(" %d", j*10)

			mutex.Lock() // 要加锁/解锁，否则 bufWriter.WriteString 写入数据有问题
			_, err := bufWriter.WriteString(strId + strName + strScore + "\n")
			if err != nil {
				fmt.Printf("WriteDataToTxt WriteString err: %v\n", err)
				return
			}
			mutex.Unlock()

			// bufWriter.Flush() // 刷入磁盘（错误示例：WriteDataToTxt err: short write，short write；因为循环太快，有时写入的数据量太小了）
		}(i)
	}
	wg.Wait()
	bufWriter.Flush() // 刷入磁盘（正确示例，io 通过 flush 操作将缓冲写入真实的文件的，所以一定要在关闭文件之前先flush，否则会造成数据丢失的情况）
}
