package main

import "fmt"

func main() {
	var i, j int = 1, 2

	// var宣言の代わり
	// :=の代入文を使って暗黙的な型宣言
	// 関数内のみで有効
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}