// package main

// import "fmt"

// type PayMethod interface {
// 	Pay(int)
// }

// type CreditCard struct {
// 	balance int
// 	limit   int
// }

// func (c *CreditCard) Pay(amount int) {
// 	if c.balance+amount > c.limit {
// 		fmt.Println("余额不足: ", c.limit-c.balance)
// 		return
// 	}
// 	c.balance += amount
// 	fmt.Println("支付成功，消费: ", amount)
// 	fmt.Println("当前余额: ", c.limit-c.balance)
// }

// func main() {
// 	c := CreditCard{balance: 100, limit: 1000}
// 	c.Pay(200)

// 	var p PayMethod = &c
// 	fmt.Println(p)

// }

package main

import (
	"fmt"
	"time"
)

// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	time.Sleep(10 * time.Second)
	close(ch)
}

func main() {
	// 创建一个带缓冲的channel
	ch := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接收goroutine
	go receiveOnly(ch)

	// 使用select进行多路复用
	timeChan := time.After(5 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case t := <-timeChan:
			fmt.Println(t)

		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
