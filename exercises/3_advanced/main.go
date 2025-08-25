package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 8.Channel
func producer(channel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}
func comsumer(channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for v := range channel {
		fmt.Println("Received ", v)
		count++
	}
	fmt.Printf("Received %d numbers from channel\n", count)
}

// 9. Lock
type Counter struct {
	mu    sync.Mutex
	count uint16
}

func (c *Counter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	//7. Channel
	channel := make(chan int)
	go func(c chan<- int) {
		for i := 1; i <= 10; i++ {
			c <- i
			fmt.Println("Send: ", i)
		}
		close(c)
	}(channel)

	go func(c <-chan int) {
		for v := range c {
			fmt.Println("Received: ", v)
		}
	}(channel)
	time.Sleep(500)

	//8. Channel
	var wg sync.WaitGroup
	wg.Add(2)
	channel_buffer := make(chan int, 10)
	go producer(channel_buffer, &wg)
	go comsumer(channel_buffer, &wg)
	wg.Wait()

	//9. Lock
	var wg2 sync.WaitGroup
	wg2.Add(10)
	counter := Counter{mu: sync.Mutex{}, count: 0}
	for i := 1; i <= 10; i++ {
		go func() {
			defer wg2.Done()
			for i := 0; i < 1000; i++ {
				counter.Add()
			}
			fmt.Printf("Goroutine %d finished count: 1000 times\n", i)
		}()
	}
	wg2.Wait()
	fmt.Println("Counter value: ", counter.count)

	//10. Lock
	var count int32
	wg2.Add(10)
	for j := 0; j < 10; j++ {
		go func() {
			defer wg2.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg2.Wait()
	fmt.Println("Count: ", count)
}
