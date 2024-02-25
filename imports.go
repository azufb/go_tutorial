package main

// 括弧でパッケージのインポートをグループ化（factoredインポートステートメント）
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}