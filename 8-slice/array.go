package main

import "fmt"

func printArray(Array [4]int) {
	for index, value := range Array {
		fmt.Println("index:", index, "value:", value)
	}
}
func main() {
	//var myArray1 [10]int
	//myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	printArray(myArray3)
}
