package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	b := decimal.NewFromFloat(1.17)
	addMoney := b.Add(b)
	fmt.Println(b, addMoney)

	a, _ := decimal.NewFromString("14.2")
	c := a.BigInt()
	fmt.Println(a, c)
}
