package main

import "fmt"

func main() {
	// バッファとして使える
	// 第2引数にバッファの長さを指定する
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}