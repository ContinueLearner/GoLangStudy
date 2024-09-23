package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("write a book")
}

func main() {
	b := &Book{}

	var r Reader = b
	r.ReadBook()

	var w Writer = r.(Writer)
	w.WriteBook()

}
