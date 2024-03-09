package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// これ以上の送信する値がないことを示すため、チャネルを close できる
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		// closeするまでループが続く（closeする必要がある）
		fmt.Println(i)
	}
}