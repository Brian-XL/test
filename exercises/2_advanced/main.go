package main

import (
	"fmt"
	"time"
)

// 1. Pointer
func pointer(num *int) {
	*num += 10
}

// 2. Pointer

func pointerSlice(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		(*slice)[i] *= 2
	}
}

// 3. Goroutine
func print() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println("奇数：", i)
			}

		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("偶数：", i)
			}
		}
	}()

	time.Sleep(1 * time.Second)
}

// 4. Goroutine
func tasksScheduler(tasks []func()) {
	for i, task := range tasks {
		go func() {
			start := time.Now()
			defer func() {
				fmt.Println("任务", i, "执行时间：", time.Since(start))
			}()
			task()
		}()
	}
	time.Sleep(5 * time.Second)
}

// 5. Obj
type Shape interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() (area float32) {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() (perimeter float32) {
	return 3.14 * 2 * c.Radius
}

type Rectangle struct {
	width  float32
	height float32
}

func (r Rectangle) Area() (area float32) {
	return r.width * r.height
}
func (r Rectangle) Perimeter() (perimeter float32) {
	return 2 * (r.width + r.height)
}

//6. obj

type Person struct {
	Name string
	Age  uint8
}
type Employee struct {
	Person
	EmployeeID uint32
}

func (e Employee) PrintInfo() {
	fmt.Printf("Employee - ID: %d  Name: %s  Age: %d\n", e.EmployeeID, e.Name, e.Age)
}

func main() {
	// 1. pointer
	a := 10
	pointer(&a)
	fmt.Println(a)
	// 1. pointer
	arr := []int{1, 2, 3, 4, 5, 99, 10}
	pointerSlice(&arr)
	fmt.Println(arr)

	// 3. goroutine
	print()

	//4. tasks scheduler
	tasks := []func(){
		func() { time.Sleep(1 * time.Second) },
		func() { time.Sleep(2 * time.Second) },
		func() { time.Sleep(3 * time.Second) },
	}
	tasksScheduler(tasks)

	// 5. obj
	c1 := Circle{Radius: 5}
	c2 := Circle{Radius: 10}
	r1 := Rectangle{width: 5, height: 2}
	r2 := Rectangle{width: 12, height: 8}

	shapes := []Shape{c1, c2, r1, r2}

	for _, s := range shapes {
		fmt.Printf("Shape %#v's Area: %.1f\n", s, s.Area())
		fmt.Printf("Shape %#v's Perimeter: %.1f\n", s, s.Perimeter())
	}
	//6. obj
	employee := []Employee{
		{Person: Person{Name: "Jack", Age: 20}, EmployeeID: 20171356},
		{Person{Name: "Alice", Age: 17}, 202422222},
		{Person{"John", 24}, 2013321313},
	}

	for _, v := range employee {
		v.PrintInfo()
	}

}
