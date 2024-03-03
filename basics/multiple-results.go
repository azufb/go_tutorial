package main

import "fmt"

// 複数の戻り値を返すこともできる
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}