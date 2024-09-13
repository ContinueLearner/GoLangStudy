package main

import "fmt"

//函数的写法
func foo1(a int,b string) int {
	fmt.Println("======foo1=========")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	c := 100
	
	return c
}

//多返回值的写法,匿名返回值
func foo2(a int ,b string)(int,int){
	fmt.Println("======foo2=========")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	return 666,777
}

//带名字的返回值，初始的时候为0
func foo3(a int,b string)(r1 int,r2 int){
	fmt.Println("======foo3=========")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	r1 = 1000
	r2 = 2000

	return 
}

//返回值同一个类型，可以写在一起
func foo4(a int,b string)(r1,r2 int){
	fmt.Println("======foo4=========")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	r1 = 1000
	r2 = 2000

	return 
}
func main(){
	c := foo1(10,"abc")
	fmt.Println("c = ",c)

	a,b := foo2(29,"def")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	a,b = foo3(29,"def")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	a,b = foo4(29,"def")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
}
