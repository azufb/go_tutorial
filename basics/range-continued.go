package main

import "fmt"

func main() {
	// 10個要素が入るint型の配列
	pow := make([]int, 10)

	// indexだけほしいので、2つめの変数を省略
	for i := range pow {
		fmt.Println(i)
		pow[i] = 1 << uint(i)
	}

	// indexは省略して値だけほしい
	// 値だけほしいときはindex省略できない、アンダーバーへ代入
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}