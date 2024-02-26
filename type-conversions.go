package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// 明示的にfloat64(別の型の変数)とすることでfloat64に型変換
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}