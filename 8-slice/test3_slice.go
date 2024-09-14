package main

import "fmt"

func main() {
	//第一种声明方式
	//slice := []int{1, 2, 3}

	//第二种声明方式
	// var slice []int
	// slice = make([]int, 3)

	//第三种声明方式：声明的同时开辟空间
	// var slice []int = make([]int, 3)

	//第四种声明方式，较为常用
	slice := make([]int, 3)

	fmt.Printf("len = %d,slice = %v\n", len(slice), slice)

	if slice == nil {
		fmt.Println("slice是一个空切片")
	} else {
		fmt.Println("slice是有空间的")
	}
}
