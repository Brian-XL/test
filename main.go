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

////2.
// package main

// import (
// 	"fmt"
// 	"time"
// )

// // 只接收channel的函数
// func receiveOnly(ch <-chan int) {
// 	for v := range ch {
// 		fmt.Printf("接收到: %d\n", v)
// 	}
// }

// // 只发送channel的函数
// func sendOnly(ch chan<- int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Printf("发送: %d\n", i)
// 	}
// 	time.Sleep(10 * time.Second)
// 	close(ch)
// }

// func main() {
// 	// 创建一个带缓冲的channel
// 	ch := make(chan int, 3)

// 	// 启动发送goroutine
// 	go sendOnly(ch)

// 	// 启动接收goroutine
// 	go receiveOnly(ch)

// 	// 使用select进行多路复用
// 	timeChan := time.After(5 * time.Second)
// 	for {
// 		select {
// 		case v, ok := <-ch:
// 			if !ok {
// 				fmt.Println("Channel已关闭")
// 				return
// 			}
// 			fmt.Printf("主goroutine接收到: %d\n", v)
// 		case t := <-timeChan:
// 			fmt.Println(t)

// 		default:
// 			fmt.Println("没有数据，等待中...")
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }

package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"

	mygorm "test/exercises/d4_database"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Age      uint8
	Birthday time.Time
}

func DBTest() {
	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connect to db failed", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("auto migrate failed", err)
	}

	users := []User{
		{Name: "John", Age: 18, Birthday: time.Now()},
		{Name: "Alice", Age: 20, Birthday: time.Now()},
		{Name: "Bob", Age: 22, Birthday: time.Now()},
	}

	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal("create users failed", result.Error)
	}
	fmt.Printf("插入成功，影响行数: %d\n", result.RowsAffected)

	fmt.Println("Print slice of users: ")
	for _, v := range users {
		fmt.Println(v)
	}
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//mysqlx.InitializeSqlx(dsn)

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//mygorm.InitializeGorm(db)
	//mygorm.SearchAssociation(db)
	//mygorm.MostCommentsPost(db)

	//mygorm.CreatePost(db)
	mygorm.DeleteComment(db)
}
