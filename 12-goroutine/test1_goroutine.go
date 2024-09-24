package main

import (
	"fmt"
	"time"
)

func newTask() {
	for i := 0; ; i++ {
		fmt.Printf("new Goroutine :i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()

	fmt.Println("main goroutine exit")

	// for i := 0; ; i++ {
	// 	fmt.Printf("main Goroutine :i = %d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }
}
