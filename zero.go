package main

import "fmt"

func main() {
	// 変数は初期値を与えずに宣言すると、ゼロ値が入る
	// 数値系＝0, bool型＝false, string型＝""(空文字)
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}