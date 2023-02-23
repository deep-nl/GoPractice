package main

import (
	"fmt"
	"math/rand"
	"time"
	// "./pointer"
)

type Player struct {
	health int
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Printf("player health: %d\r", p.health)
		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		p.health -= rand.Intn(40)
		if p.health <= 0 {
			fmt.Println("\nGameOver")
			break
		}
		<-ticker.C
	}
}

func main() {
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}

// func main() {
//    pointer.Test()
//    /* 定义局部变量 */
//    var a int = 10

//    /* 循环 */
//    LOOP: for a < 20 {
//       if a == 15 {
//          /* 跳过迭代 */
//          a = a + 2
//          goto LOOP
//       }
//       fmt.Printf("a的值为 : %d\n", a)
//     //   fmt.Println("a的值为 : %d", a)
//       a++
//    }
// }
