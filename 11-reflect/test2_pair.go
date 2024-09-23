package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader = tty

	var w io.Writer = r.(io.Writer)

	w.Write([]byte("Hello THIS is a test\n"))

}
