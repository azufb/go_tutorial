package main

import (
	"fmt"
	"math"
)

func main() {
	// 大文字で始まる名前＝外部のパッケージから参照できる、exportされている
	// 小文字で始まる名前＝外部から参照できない、exportされていない
	fmt.Println(math.Pi)
}