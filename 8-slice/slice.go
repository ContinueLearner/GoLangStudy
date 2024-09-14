package main

import "fmt"

func printArray1(array []int) {
	for _, value := range array {
		fmt.Println("value:", value)
	}
}
func main() {
	myArray := []int{1, 2, 3, 4}
	printArray1(myArray)
}
