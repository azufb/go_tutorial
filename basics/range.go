package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// forループでスライスやマップを繰り返し処理するときにはrange
	// 反復ごとに2つの変数を返す（indexとindexの場所の要素のコピー）
	for i, v := range pow {
		fmt.Println(i)
		fmt.Printf("2**%d = %d\n", i, v)
	}
}