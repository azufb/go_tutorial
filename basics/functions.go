package main

import "fmt"

// 加算
// 変数名の後ろに型
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}