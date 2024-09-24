package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine正在运行")

		channel <- 666
	}()

	num := <-channel

	fmt.Println("num = ", num)
	fmt.Println("main routine运行结束")
}
