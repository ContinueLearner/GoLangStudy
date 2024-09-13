package main

import "fmt"

func main(){
	var a int
	fmt.Println("a = ",a)
	fmt.Printf("type of a = %T\n",a)

	var b int = 10
	
	fmt.Println("b = ",b)
	fmt.Printf("type of b = %T\n",b)

	var c = 3.14

	fmt.Println("c = ",c)
	fmt.Printf("type of c = %T\n",c)

	d := "abcd"
	fmt.Println("d = ",d)
	fmt.Printf("type of d = %T\n",d)

	var aa,bb = 100,3.14
	fmt.Println("aa = ",aa)
	fmt.Println("bb = ",bb)

	var(
		xx = "abcd"
		yy = true
	)
	fmt.Println("xx = ",xx)
	fmt.Println("yy = ",yy)
}