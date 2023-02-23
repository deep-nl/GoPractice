package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
	基本需求

	func main() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

		// 模拟并发进行的处理业务逻辑
		for i := 1; i < 10; i++ {
			go func(i int) {
				for {
					// 我们希望程序能等当前这个周期休眠完，再优雅退出
					time.Sleep(time.Duration(i) * time.Second)
					fmt.Printf("gorotine %d end\n", i)
				}
			}(i)
		}
		//fmt.Println(<-sig)
		//fmt.Println("main goroutine end")
	}
*/
//func main() {
//	signal_1()
//}

// 方案一：增加两个channel
func signal_1() {
	sig := make(chan os.Signal)
	stopCh := make(chan struct{})
	finishedCh := make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT)

	go func(stopCh, finishedCh chan struct{}) {
		for {
			select {
			case <-stopCh:
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(3 * time.Second)
				fmt.Println("Sleep 3 seconds")
			}
		}
	}(stopCh, finishedCh)

	<-sig
	stopCh <- struct{}{}
	<-finishedCh
	fmt.Println("finished")
}

// 方案二，channel嵌套
func main() {
	sig := make(chan os.Signal)
	stopCh := make(chan chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopChh chan chan struct{}) {
		for {
			select {
			case ch := <-stopCh:
				// 结束后，通过ch通知主goroutine
				fmt.Println("stopped")
				ch <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh)

	<-sig
	// ch作为一个channel，传递给子goroutine，待其结束后从中返回
	ch := make(chan struct{})
	stopCh <- ch
	<-ch
	fmt.Println("finished")
}
