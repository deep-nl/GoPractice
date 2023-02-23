package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := big.NewInt(100)
	j := big.NewInt(200)
	k := i.Add(i, j)
	fmt.Println(i, j, k)
	i.Add(i, j)
	// 虽然这里没有了k，但是还是会影响k的值
	fmt.Println(i, j, k)
	// 如何操作，看下面
	i = big.NewInt(0).Add(i, j)
	fmt.Println(i, j, k)
	string2big()
}

func string2big() {
	s := big.NewInt(0)
	str := "Hello the fucking world"
	fmt.Println(str)
	s.SetString(str, 0) // 不支持直接转换
	s.SetBytes([]byte(str))
	fmt.Println(s)
	i := big.NewInt(0)
	i = i.Add(i, s)
	fmt.Println(i)

}
