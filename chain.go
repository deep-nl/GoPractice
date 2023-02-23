package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
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
func chain1() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	fmt.Println(<-sig)
}

//func main() {
//	stop := make(chan bool)
//
//	go func() {
//		for {
//			select {
//			case <-stop:
//				fmt.Println("监控退出，停止了...")
//				return
//			default:
//				fmt.Println("goroutine监控中...")
//				time.Sleep(2 * time.Second)
//			}
//		}
//	}()
//
//	time.Sleep(10 * time.Second)
//	fmt.Println("可以了，通知监控停止")
//	stop <- true
//	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
//	time.Sleep(5 * time.Second)
//}
