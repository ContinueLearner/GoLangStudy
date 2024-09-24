package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 3)

	go func() {
		defer fmt.Println("goroutine结束")

		for i := 0; i < 3; i++ {
			channel <- i
			fmt.Println("子go程正在运行, 发送的元素=", i, " len(c)=", len(channel), ", cap(c)=", cap(channel))
		}
	}()

	time.Sleep(time.Second * 1)

	for i := 0; i < 3; i++ {
		num := <-channel
		fmt.Println("num = ", num)
	}

	fmt.Println("mainroutine结束")
}
