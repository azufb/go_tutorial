package main

import (
	"fmt"
	"math"
)

// 関数も変数
// 関数を関数の引数に渡すことができ、戻り値としても使える
func compute(fn func(float64, float64) float64) float64 {
	// 引数として受け取った関数に引数を渡して実行
	// 関数の結果を返す
	return fn(3, 4)
}

func main() {
	// 関数を変数hypotに代入
	hypot := func(x, y float64) float64 {
		// 平方根を返す
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
