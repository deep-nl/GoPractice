package main

/*
阻塞主协程的几个方法

This method is not recommended, as it can cause the program to become unresponsive,
and it is better to use other methods like channels, mutexes, or semaphores to synchronize the behavior of multiple goroutines.
Additionally, using an infinite loop to hang the main goroutine is not a clean or graceful way to handle the situation,
and it can make it more difficult to understand the behavior of the program.
*/
import (
	"fmt"
	"time"
)

func main() {
	main1()
}

// using an empty select statement to block the program
func main1() {
	go func() {
		fmt.Println("Starting a goroutine")
		time.Sleep(3 * time.Second)
		fmt.Println("Goroutine finished")
	}()

	fmt.Println("Main function starting")
	select {}
	fmt.Println("Main function finished")
}

// more graceful
func main2() {
	done := make(chan bool)

	go func() {
		fmt.Println("Starting a goroutine")
		time.Sleep(3 * time.Second)
		fmt.Println("Goroutine finished")
		done <- true
	}()

	fmt.Println("Main function starting")
	<-done
	fmt.Println("Main function finished")
}

// 类似于rust的loop{}
func main3() {
	go func() {
		fmt.Println("Starting a goroutine")
		time.Sleep(3 * time.Second)
		fmt.Println("Goroutine finished")
	}()

	fmt.Println("Main function starting")
	for {
	}
}
