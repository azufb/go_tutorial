package main

import "fmt"

func main() {
	m := make(map[string]int)

	// 要素の挿入
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// 要素の更新
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 要素の取得と代入
	elm := m["Answer"]
	fmt.Println(elm)
	elm = 100
	fmt.Println(elm)

	// 要素の削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 要素が存在するかどうかを2つ目の値で確認できる
	// true/falseで返ってくる
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 500
	v2, ok2 := m["Answer"]
	fmt.Println("The value:", v2, "Present?", ok2)
	
	fmt.Println(m)
}