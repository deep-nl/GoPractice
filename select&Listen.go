package main

import "time"

func main() {
	ch := make(chan int)
	select {

	case i := <-ch:
		println(i)

	default:
		println("start")
		time.Sleep(2 * time.Second)
	}
	// 加上for可以持续监听
	for {
		select {
		case i := <-ch:
			println(i)

		default:
			println("listening")
			time.Sleep(2 * time.Second)
		}
	}
}
