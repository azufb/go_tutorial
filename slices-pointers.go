package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// 配列namesの0番目と1番目の値を要素とするスライス
	a := names[0:2]
	// 配列namesの1番目と2番目の値を要素とするスライス
	b := names[1:3]
	fmt.Println(a, b)

	// スライスには元の配列の参照が入っているため、
	// 変更すると、そのメモリアドレスの値が変わることになり、元の配列の値まで変わってしまう
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}