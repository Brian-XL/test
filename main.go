package main

import "fmt"

type Printer interface {
	Print()
}

type Book struct {
	name string
}

func (b Book) Print() {
	fmt.Println("Book:", b.name)
}

func main() {
	var p Printer = Book{"Golang"}

	p.Print()

	b, ok := p.(Book) //type assertion

	if ok {
		fmt.Println("Book:", b.name)
	} else {
		fmt.Println("Not a book")
	}
}
