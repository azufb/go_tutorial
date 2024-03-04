package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append(配列名, 追加する要素)で配列に要素を追加できる
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	// 一度に複数の要素を追加することもできる
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}