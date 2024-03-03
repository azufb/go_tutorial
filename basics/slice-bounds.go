package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	// 下限を省略
	// この場合、スライスsの1番目と2番目の要素
	s = s[:2]
	fmt.Println(s)

	// 上限を省略
	// この場合、スライスの3番目の要素だけ
	s = s[1:]
	fmt.Println(s)
}