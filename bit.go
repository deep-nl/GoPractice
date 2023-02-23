package main

import (
	"fmt"
	"go-base/logging"
)

/*
&      位运算 AND
|      位运算 OR
^      位运算 XOR
&^     位清空 (AND NOT)
<<     左移
>>     右移
*/

// 在使用const关键字的时候 可使用内置变量 iota从0开始自动递增
// 在遇到下个常量块或者单个常量定义的时候 也就是再一次使用const关键字的时候 iota置0
type ByteSize float64

const (
	B  ByteSize = 1 << (10 * iota) // 1<<(10*0)
	KB                             // 1<<(10*1) 左移动10位  2的10次方=1024
	MB                             // 1<<(10*2)
	GB                             // 1<<(10*3)
	TB                             // 1<<(10*4)
	PB                             //  1<<(10*5)
)

func main() {
	t := logging.Logger.WithField("name", "bit compute")
	// >> <<
	var i int64
	t.Info(fmt.Sprintf("%b", i)) //0
	i = 1
	t.Info(fmt.Sprintf("%b %d", i, i))       // 1 1
	t.Info(fmt.Sprintf("%b %d", i<<1, i<<1)) //10 2  左移动 每个移位位置代表2的幂,左移动增加
	t.Info(fmt.Sprintf("%b %d", i<<2, i<<2)) // 100 4

	//
	t.Info(fmt.Sprintf("%v %v %v %v %v %v", B, KB, MB, GB, TB)) // 100 4

	// 位运算 & AND
	c, d := 1, 1
	t.Infof("c&d %b %b %b", c&d, c, d) //c&d 1 1 1
	c, d = 1, 0
	t.Infof("c&d %b %b %b", c&d, c, d) //c&d 0 1 0
	c, d = 5, 3
	t.Infof("c&d %b %b %b", c&d, c, d) //c&d 1 101 011
	c, d = 5, 4
	t.Infof("c&d %b %b %b", c&d, c, d) //c&d 100 101 100  十进制4
	c, d = 50, 43
	t.Infof("c&d %b %b %b", c&d, c, d) //c&d 100010 110010 101011 十进制34

	ret := addWithoutOperator(c, d)
	t.Warnln(ret)
	/*
	 a |= b  ----->  a = a | b  , a 或者 b 只要有一个为 1, 那么，a 的最终结果就为 1
	 a &= b  ----->  a = a & b  , a 和 b 二者必须都为 1, 那么，a 的最终结果才为 1
	 a ^= b  ----->  a = a ^ b  , 当且仅当 a 和 b 的值不一致时，a 的最终结果才为1，否则为0
	*/
	a, b := 1, 0
	a |= b
	t.Infof("%b %b", a, b) // 1 0
	a, b = 1, 0
	a &= b
	t.Infof("%b %b", a, b) // 0 0
	a, b = 1, 0
	a ^= b
	t.Infof("%b %b", a, b) //  1 0
}

func addWithoutOperator(num1 int, num2 int) int {
	a := num1 ^ num2 //如果无进位，则a就是结果
	b := num1 & num2 //记录进位，即有没有最高位1+1的情况
	b = b << 1       //进位要左移一位
	if b == 0 {
		return a
	}
	return addWithoutOperator(a, b)
}
