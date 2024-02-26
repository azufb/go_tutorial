package main

import (
	"fmt"
)

// 平方根の実装
func Sqrt(x float64) float64 {
	// 10回計算を繰り返す
	for z := 1; 1 <= 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)

		if
	}
}

func main() {
	fmt.Println(Sqrt(2))
}