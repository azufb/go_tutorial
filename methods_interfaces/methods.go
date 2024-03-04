package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// goにクラスはない
// メソッド定義
// func (レシーバ) メソッド名
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// レシーバとは
// 構造体のフィールドにアクセスする方法を提供するメソッドの引数
// 関数の引数と違って、あえて渡す必要がない

type Profile struct {
	name string
	age int
}

func (s Profile) Str() string {
	return s.name + "さん、こんにちは！"
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	s := Profile{"A", 26}
	fmt.Println(s.Str())
}