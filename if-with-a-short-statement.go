package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// 条件(v < lim)の前にステートメント
	// 条件で使う、変数vを宣言
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20), // カンマ忘れない
	)
}