package main

import "fmt"

func main() {
	// メソッドが0のインターフェースは、空のインターフェースと呼ばれる
	// 任意の型の値を保持できる
	// 未知の型の値を扱う場合に使う
	var i interface {}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}