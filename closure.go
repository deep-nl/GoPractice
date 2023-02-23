package main

import "fmt"

/*
	go语言支持函数式编程：
		支持将一个函数作为另一个函数的参数，
		也支持将一个函数作为另一个函数的返回值。

	闭包(closure)：
		一个外层函数中，有内层函数，该内层函数中，
		会操作外层函数的局部变量(外层函数中的参数，或者外层函数中直接定义的变量)，
		并且该外层函数的返回值就是这个内层函数。

		这个内层函数和外层函数的局部变量，统称为闭包结构。
		局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
		但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。
*/

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		fmt.Printf("sum地址: %p\t\t", &sum)
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}

	aa := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, aa = aa(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
