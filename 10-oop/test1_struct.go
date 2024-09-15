package main

import "fmt"

//声明一种数据类型，类似于C++中的typedef
type myInt int

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	book.auth = "lisi"
}

func changeBook2(book *Book) {
	book.auth = "lisi"
}

func main() {
	var book1 Book
	book1.auth = "zhangsan"
	book1.title = "huozhe"

	fmt.Printf("%v", book1)

	changeBook(book1)
	fmt.Printf("%v", book1)

	changeBook2(&book1)
	fmt.Printf("%v", book1)

}
