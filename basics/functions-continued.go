package main

import "fmt"

// 2つ以上の引数が同じ型だったら、まとめて書ける
// 引数１, 引数２ 型
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}