package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"zhangs"}

	myFunc(book)
	myFunc(100)
	myFunc(3.14)
	myFunc("abc")

}
