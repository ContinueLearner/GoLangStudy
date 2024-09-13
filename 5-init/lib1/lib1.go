package lib1

import "fmt"

func Lib1Test(){
	fmt.Println("Lib1Test 调用")
}
func init(){
	fmt.Println("lib1.go init()")
}