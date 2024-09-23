package main

import "fmt"

func main() {
	var a string = "lipengyu"

	var allType interface{} = a

	str, _ := allType.(string)

	fmt.Println(str)

}
