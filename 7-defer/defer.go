package main

import "fmt"

//defer和return一起存在时，return后面的函数会先被调用
func defer_func() int {
	fmt.Println("defer_func called")
	return 0
}
func return_func() int {
	fmt.Println("return_func called")
	return 0
}

func defer_return() int {
	defer defer_func()
	return return_func()
}
func main() {
	//函数销毁的时候defer才会被调用
	defer fmt.Println("类似于C++的析构函数")
	fmt.Println("defer.go")

	defer_return()
}
