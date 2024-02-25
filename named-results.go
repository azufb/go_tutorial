package main

import "fmt"

// 戻り値の変数に名前を付けれる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum -x

	// 名前を付けた戻り値の変数を使うとreturnには何も書かなくてよい
	// "naked" return
	return
}

func main() {
	fmt.Println(split(17))
}