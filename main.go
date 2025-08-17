package main

import "fmt"

type PayMethod interface {
	Pay(int)
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) {
	if c.balance+amount > c.limit {
		fmt.Println("余额不足: ", c.limit-c.balance)
		return
	}
	c.balance += amount
	fmt.Println("支付成功，消费: ", amount)
	fmt.Println("当前余额: ", c.limit-c.balance)
}

func main() {
	c := CreditCard{balance: 100, limit: 1000}
	c.Pay(200)

	var p PayMethod = &c
	fmt.Println(p)

}
