package main

import "fmt"

//const主要是和iota用作枚举使用

const (
	a,b = iota+1,iota+2
	c,d
	e,f

	g,h = iota * 2,iota * 3
	i,k
)
func main(){
	//该值不能修改
	const length int = 10
	fmt.Println("length = ",length)

	//length = 100

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)
}