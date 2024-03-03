package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 初期化子がある場合は、型を省略できる
// 初期化子が持つ型になる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}