package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)

	numbers = append(numbers, 1)
	numbers = append(numbers, 1)
	numbers = append(numbers, 1)

	fmt.Printf("len = %d,cap = %d,slice = %v", len(numbers), cap(numbers), numbers)

}
