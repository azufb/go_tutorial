package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// スライスの長さを0にするようスライス
	s = s[:0]
	printSlice(s)

	// スライスの長さを0から4に増やす
	s = s[:4]
	printSlice(s)

	// スライスの最初の2つの要素は省く
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}